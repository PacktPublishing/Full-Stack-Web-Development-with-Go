package main

import (
	"chapter6/gen"
	gen "chapter6/gen"
	"chapter6/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type server struct {
	router         *mux.Router
	database       *gen.Queries
	hListSetsFn    http.HandlerFunc
	hSaveSetsFn    http.HandlerFunc
	hListExercises http.HandlerFunc
	hLogin         http.HandlerFunc
	hSaveExercise  http.HandlerFunc
}

func NewServer(db *chapter6.Queries) server {
	s := server{
		database: db,
		router:   new(mux.Router),
	}

	return s
}

func (s *server) SetupHandlerFunction() {
	m := model.Sets{}
	s.hListSetsFn = handlerListSets(m, s.database)
	s.hSaveSetsFn = handlerSaveSets(m, s.database)

	e := model.Exercise{}
	s.hListExercises = handleListExercises(e, s.database)
	s.hSaveExercise = handleSaveExercise(e, s.database)

	l := model.Login{}
	s.hLogin = handleLogin(l, s.database)
}

func basicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		log.Println("Middleware called on", req.URL.Path)
		// do stuff
		h.ServeHTTP(wr, req)
	})
}

func handlerListSets(m model.SetsInterface, database *gen.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := m.ListSets(database)

		if d == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{Msg: "error getting sets data"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(d)
	}
}

func handlerSaveSets(m model.SetsInterface, database *gen.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = m.AddSets(database)

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{Msg: "error inserting new sets"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.ExerciseResponse{Success: true})
	}

}

func handleListExercises(m model.Exercise, database *gen.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := m.ListExercises(database)

		if d == nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{Msg: "error getting exercise data"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(d)
	}
}

func handleLogin(lg model.Login, database *gen.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s = http.StatusOK
		var l = model.LoginResponse{}

		err := json.NewDecoder(r.Body).Decode(&lg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if lg.ValidateUser(database) {
			l.Success = true
		} else {
			s = http.StatusBadRequest
			l.Success = false
		}

		w.WriteHeader(s)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(l)
	}
}

func handleSaveExercise(m model.Exercise, database *gen.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = m.AddExercise(database)

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.ErrorResponse{Msg: "error inserting new exercise"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(model.ExerciseResponse{Success: true})
	}
}

func (s server) SetupRoutes() {
	//POST handler for /login
	s.router.HandleFunc("/login", s.hLogin).Methods("POST")

	//POST handler for /exercise
	s.router.HandleFunc("/exercise", s.hSaveExercise).Methods("POST")

	//GET handler for /listexercise
	s.router.HandleFunc("/listexercise", s.hListExercises)

	//POST handler for /sets
	s.router.HandleFunc("/sets", s.hSaveSetsFn).Methods("POST")

	//POST handler for /listsets
	s.router.HandleFunc("/listsets", s.hListSetsFn)

	//root will redirect to /apo
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app", http.StatusPermanentRedirect)
	})

	// Use our basicMiddleware
	s.router.Use(basicMiddleware)
}

func (s server) Run() {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         "127.0.0.1:3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
