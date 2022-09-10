package main

import (
	"chapter10/internal"
	"chapter10/internal/api"
	"chapter10/internal/env"
	"chapter10/store"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		env.GetAsString("DB_USER", "local"),
		env.GetAsString("DB_PASSWORD", "asecurepassword"),
		env.GetAsString("DB_HOST", "localhost"),
		env.GetAsInt("DB_PORT", 5432),
		env.GetAsString("DB_NAME", "fullstackdb"),
	)

	// Open the database
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatalln("Error opening database:", err)
	}

	// Connectivity check
	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping:", err)
	}

	// Create our demo user
	createUserInDb(db)

	// Start our server
	server := api.NewServer(env.GetAsInt("SERVER_PORT", 9002))

	server.MustStart()
	defer server.Stop()

	defaultMiddleware := []mux.MiddlewareFunc{
		api.JSONMiddleware,
		api.CORSMiddleware(env.GetAsSlice("CORS_WHITELIST",
			[]string{
				"http://localhost:9000",
				"http://0.0.0.0:9000",
			}, ","),
		),
	}

	// Handlers
	server.AddRoute("/login", handleLogin(db), http.MethodPost, defaultMiddleware...)
	server.AddRoute("/logout", handleLogout(), http.MethodGet, defaultMiddleware...)

	// Our session protected middleware
	protectedMiddleware := append(defaultMiddleware, validCookieMiddleware(db))
	server.AddRoute("/checkSecret", checkSecret(db), http.MethodGet, protectedMiddleware...)

	// Workouts
	server.AddRoute("/workout", handlecreateNewWorkout(db), http.MethodPost, protectedMiddleware...)
	server.AddRoute("/workout", handleListWorkouts(db), http.MethodGet, protectedMiddleware...)
	server.AddRoute("/workout/{workout_id}", handleDeleteWorkout(db), http.MethodDelete, protectedMiddleware...)
	server.AddRoute("/workout/{workout_id}", handleAddSet(db), http.MethodPost, protectedMiddleware...)
	server.AddRoute("/workout/{workout_id}/{set_id}", handleUpdateSet(db), http.MethodPut, protectedMiddleware...)

	// Wait for CTRL-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	// We block here until a CTRL-C / SigInt is received
	// Once received, we exit and the server is cleaned up
	<-sigChan
}

func createUserInDb(db *sql.DB) {

	ctx := context.Background()
	querier := store.New(db)

	log.Println("Creating user@user...")
	hashPwd := internal.HashPassword("password")

	_, err := querier.CreateUsers(ctx, store.CreateUsersParams{
		UserName:     "user@user",
		PasswordHash: hashPwd,
		Name:         "Dummy user",
	})

	// This is interesting to look at, the sql/pq library recommends we use
	// this pattern to understand errors. We could use the ErrorCode directly
	// or look for the specific type. We know we'll be violating unique_violation
	// if our user already exists in the database
	if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
		log.Println("Dummy User already present")
		return
	}

	if err != nil {
		log.Println("Failed to create user:", err)
	}
}
