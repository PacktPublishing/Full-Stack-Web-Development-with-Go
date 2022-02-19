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

	// Create the store
	store := store.New(db)
	// Create our demo user
	createUserDb(store)

	// Start our server
	server := api.NewServer(internal.GetAsInt("SERVER_PORT", 9002))

	if err := server.Start(); err != nil {
		log.Fatalln("can't start API server", err)
	}
	defer server.Stop()

	protectedMiddleware := []mux.MiddlewareFunc{
		api.JSONMiddleware,
	}

	server.AddRoute("/test", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Here")
	}, http.MethodGet, protectedMiddleware...)

	// Wait for CTRL-C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	// We block here until a CTRL-C / SigInt is received
	// Once received, we exit and the server is cleaned up
	<-sigChan
}

func createUserDb(s *store.Queries) {
	//has the user been created
	c := context.Background()
	u, _ := s.GetUserByName(c, "user@user")

	if u.UserName == "user@user" {
		log.Println("user@user exist...")
		return
	}
	log.Println("Creating user@user...")
	hashPwd, _ := internal.HashPassword("password")
	_, err := s.CreateUsers(c, store.CreateUsersParams{
		UserName:     "user@user",
		PasswordHash: hashPwd,
		Name:         "Dummy user",
	})
	if err != nil {
		log.Printf("error creating test user: user@user - failed with err %v", err)
	}
}
