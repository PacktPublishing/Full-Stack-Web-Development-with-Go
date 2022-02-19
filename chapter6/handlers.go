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

type loginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func handleLogin(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {

		// Thanks to our middleware, we know we have JSON
		// we'll decode it into our request type and see if it's valid
		payload := loginRequest{}
		if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
			log.Println("Error decoding the body", err)
			api.JSONError(wr, http.StatusBadRequest, "Error decoding JSON")
			return
		}

		querier := store.New(db)
		user, err := querier.GetUserByName(req.Context(), payload.Username)
		if errors.Is(err, sql.ErrNoRows) || !internal.CheckPasswordHash(payload.Password, user.PasswordHash) {
			api.JSONError(wr, http.StatusForbidden, "Bad Credentials")
			return
		}
		if err != nil {
			log.Println("Received error looking up user", err)
			api.JSONError(wr, http.StatusInternalServerError, "Couldn't log you in due to a server error")
			return
		}

		// We're valid. Let's tell the user and set a cookie
		session, err := cookieStore.Get(req, "session-name")
		if err != nil {
			log.Println("Cookie store failed with", err)
			api.JSONError(wr, http.StatusInternalServerError, "Session Error")
		}
		session.Values["userAuthenticated"] = true
		session.Values["userID"] = user.UserID
		session.Save(req, wr)
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
		_ = json.NewEncoder(wr).Encode(struct {
			Message string
		}{Message: fmt.Sprintf("Hello there %s", user.Name)})
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
			log.Println("Failed to save session")
			api.JSONError(wr, http.StatusInternalServerError, "Session Error")
			return
		}

		_ = json.NewEncoder(wr).Encode(struct {
			Message string
		}{Message: "logout ok"})
	})
}
