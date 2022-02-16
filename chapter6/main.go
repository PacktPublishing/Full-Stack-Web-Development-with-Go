package main

import (
	gen "chapter6/gen"
	"chapter6/model"
	"chapter6/pkg"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	dbQuery *gen.Queries
	store   = sessions.NewCookieStore([]byte("forDemo"))
)

//loginHandler handles authentication
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var lg model.Login
	var s = http.StatusOK
	var l = model.LoginResponse{}

	err := json.NewDecoder(r.Body).Decode(&lg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if lg.ValidateUser(dbQuery) {
		l.Success = true
	} else {
		s = http.StatusBadRequest
		l.Success = false
	}

	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(l)
}

func basicMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		log.Println("Middleware called on", req.URL.Path)
		// do stuff
		h.ServeHTTP(wr, req)
	})
}

func main() {
	initDatabase()

	router := mux.NewRouter()

	//POST handler for /login
	router.HandleFunc("/login", loginHandler).Methods("POST")

	//POST handler for /exercise
	router.HandleFunc("/exercise", postExerciseHandler).Methods("POST")

	//GET handler for /listexercise
	router.HandleFunc("/listexercise", listExerciseHandler)

	//POST handler for /sets
	router.HandleFunc("/sets", postSetsHandler).Methods("POST")

	//POST handler for /listsets
	router.HandleFunc("/listsets", listSetsHandler)

	//root will redirect to /apo
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/app", http.StatusPermanentRedirect)
	})

	// Use our basicMiddleware
	router.Use(basicMiddleware)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func listSetsHandler(writer http.ResponseWriter, request *http.Request) {

}

func postSetsHandler(w http.ResponseWriter, r *http.Request) {
	var s model.Sets

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.AddSets(dbQuery)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Msg: "error inserting new sets"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ExerciseResponse{Success: true})
}

func listExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var e model.Exercise

	l := e.ListExercises(dbQuery)

	if l == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Msg: "error getting exercise data"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(l)
}

func postExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var e model.Exercise

	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = e.AddExercise(dbQuery)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Msg: "error inserting new exercise"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.ExerciseResponse{Success: true})
}

func initDatabase() {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		GetAsString("DB_USER", "postgres"),
		GetAsString("DB_PASSWORD", "mysecretpassword"),
		GetAsString("DB_HOST", "localhost"),
		GetAsInt("DB_PORT", 5432),
		GetAsString("DB_NAME", "postgres"),
	)

	// Open the database
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	// Connectivity check
	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping:", err)
	}

	// Create the store
	dbQuery = gen.New(db)

	ctx := context.Background()

	createUserDb(ctx)

	if err != nil {
		os.Exit(1)
	}
}

func createUserDb(ctx context.Context) {
	//has the user been created
	u, _ := dbQuery.GetUserByName(ctx, "user@user")

	if u.UserName == "user@user" {
		log.Println("user@user exist...")
		return
	}
	log.Println("Creating user@user...")
	hashPwd, _ := pkg.HashPassword("password")
	_, err := dbQuery.CreateUsers(ctx, gen.CreateUsersParams{
		UserName:     "user@user",
		PassWordHash: hashPwd,
		Name:         "Dummy user",
	})
	if err != nil {
		log.Println("error getting user@dummyuser.domain ", err)
		os.Exit(1)
	}
}
