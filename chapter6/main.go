package main

import (
	"chapter6/api"
	"chapter6/internal"
	"chapter6/store"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var (
	cookieStore = sessions.NewCookieStore([]byte("forDemo"))
)

func main() {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		internal.GetAsString("DB_USER", "local"),
		internal.GetAsString("DB_PASSWORD", "asecurepassword"),
		internal.GetAsString("DB_HOST", "localhost"),
		internal.GetAsInt("DB_PORT", 5432),
		internal.GetAsString("DB_NAME", "fullstackdb"),
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
	server := api.NewServer(internal.GetAsInt("SERVER_PORT", 9002))

	if err := server.Start(); err != nil {
		log.Fatalln("can't start API server", err)
	}
	defer server.Stop()

	defaultMiddleware := []mux.MiddlewareFunc{
		api.JSONMiddleware,
	}

	// Handlers
	server.AddRoute("/login", handleLogin(db), http.MethodPost, defaultMiddleware...)
	server.AddRoute("/logout", handleLogout(), http.MethodGet, defaultMiddleware...)

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
	hashPwd, _ := internal.HashPassword("password")

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
