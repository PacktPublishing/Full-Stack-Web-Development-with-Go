package main

import (
	"chapter6/internal"
	"chapter6/internal/api"
	"chapter6/store"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	cookieStore = sessions.NewCookieStore([]byte("forDemo"))
)

func init() {
	cookieStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 15,
		HttpOnly: true,
	}
}

type LoginInterface interface {
	GetUserByName(req *http.Request, wr http.ResponseWriter, q *store.Queries, payload LoginRequest)
}

type Login struct {
}

func (l Login) GetUserByName(req *http.Request, wr http.ResponseWriter, q *store.Queries, payload LoginRequest) {
	user, err := q.GetUserByName(req.Context(), payload.Username)
	if errors.Is(err, sql.ErrNoRows) || !internal.CheckPasswordHash(payload.Password,
		user.PasswordHash) {
		api.JSONError(wr, http.StatusForbidden, "Bad Credentials")
		return
	}
	if err != nil {
		log.Println("Received error looking up user", err)
		api.JSONError(wr, http.StatusInternalServerError, "Couldn't log you in due to a server error")
		return
	}

	session, err := cookieStore.Get(req, "session-name")
	if err != nil {
		log.Println("Cookie store failed with", err)
		api.JSONError(wr, http.StatusInternalServerError, "Session Error")
	}
	session.Values["userAuthenticated"] = true
	session.Values["userID"] = user.UserID
	session.Save(req, wr)
}

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func handleLogin(li LoginInterface, querier *store.Queries) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {

		payload := LoginRequest{}
		if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
			log.Println("Error decoding the body", err)
			api.JSONError(wr, http.StatusBadRequest, "Error decoding JSON")
			return
		}

		li.GetUserByName(req, wr, querier, payload)
	})
}

func checkSecret(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		userDetails, _ := userFromSession(req)

		querier := store.New(db)
		user, err := querier.GetUser(req.Context(), userDetails.UserID)
		if errors.Is(err, sql.ErrNoRows) {
			api.JSONError(wr, http.StatusForbidden, "User not found")
			return
		}

		api.JSONMessage(wr, http.StatusOK, fmt.Sprintf("Hello there %s", user.Name))
	})
}

func handleLogout() http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		session, err := cookieStore.Get(req, "session-name")
		if err != nil {
			log.Println("Cookie store failed with", err)
			api.JSONError(wr, http.StatusInternalServerError, "Session Error")
			return
		}

		session.Options.MaxAge = -1 // deletes
		session.Values["userID"] = int64(-1)
		session.Values["userAuthenticated"] = false

		err = session.Save(req, wr)
		if err != nil {
			api.JSONError(wr, http.StatusInternalServerError, "Session Error")
			return
		}

		api.JSONMessage(wr, http.StatusOK, "logout successful")
	})
}

func handlecreateNewWorkout(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		userDetails, ok := userFromSession(req)
		if !ok {
			api.JSONError(wr, http.StatusForbidden, "Bad context")
			return
		}
		querier := store.New(db)

		res, err := querier.CreateUserWorkout(req.Context(), userDetails.UserID)
		if err != nil {
			api.JSONError(wr, http.StatusInternalServerError, err.Error())
			return
		}

		json.NewEncoder(wr).Encode(&res)

	})
}

func handleListWorkouts(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		userDetails, ok := userFromSession(req)
		if !ok {
			api.JSONError(wr, http.StatusForbidden, "Bad context")
			return
		}

		querier := store.New(db)
		workouts, err := querier.GetWorkoutsForUserID(req.Context(), userDetails.UserID)
		if err != nil {
			api.JSONError(wr, http.StatusInternalServerError, err.Error())
			return
		}
		json.NewEncoder(wr).Encode(&workouts)
	})
}

func handleAddSet(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {

		workoutID, err := strconv.Atoi(mux.Vars(req)["workout_id"])
		if err != nil {
			api.JSONError(wr, http.StatusBadRequest, "Bad workout_id")
			return
		}

		type newSetRequest struct {
			ExerciseName string `json:"exercise_name,omitempty"`
			Weight       int    `json:"weight,omitempty"`
		}

		payload := newSetRequest{}
		if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
			log.Println("Error decoding the body", err)
			api.JSONError(wr, http.StatusBadRequest, "Error decoding JSON")
			return
		}

		querier := store.New(db)

		set, err := querier.CreateDefaultSetForExercise(req.Context(),
			store.CreateDefaultSetForExerciseParams{
				WorkoutID:    int64(workoutID),
				ExerciseName: payload.ExerciseName,
				Weight:       int32(payload.Weight),
			})
		if err != nil {
			api.JSONError(wr, http.StatusInternalServerError, err.Error())
			return
		}
		json.NewEncoder(wr).Encode(&set)
	})
}

func handleUpdateSet(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		// TODO
	})
}

func handleDeleteWorkout(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		userDetails, ok := userFromSession(req)
		if !ok {
			api.JSONError(wr, http.StatusForbidden, "Bad context")
			return
		}

		workoutID, err := strconv.Atoi(mux.Vars(req)["workout_id"])
		if err != nil {
			api.JSONError(wr, http.StatusBadRequest, "Bad workout_id")
			return
		}

		err = store.New(db).DeleteWorkoutByIDForUser(req.Context(), store.DeleteWorkoutByIDForUserParams{
			UserID:    userDetails.UserID,
			WorkoutID: int64(workoutID),
		})

		if err != nil {
			api.JSONError(wr, http.StatusBadRequest, "Bad workout_id")
			return
		}

		api.JSONMessage(wr, http.StatusOK, fmt.Sprintf("Workout %d is deleted", workoutID))
	})
}
