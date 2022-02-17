package main

import (
	gen "chapter6/gen"
	"chapter6/pkg"
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	data  *gen.Queries
	store = sessions.NewCookieStore([]byte("forDemo"))
)

func main() {
	initDatabase()
	s := NewServer(data)
	s.SetupHandlerFunction()
	s.SetupRoutes()
	s.Run()
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
	data = gen.New(db)

	ctx := context.Background()

	createUserDb(ctx)

	if err != nil {
		os.Exit(1)
	}
}

func createUserDb(ctx context.Context) {
	//has the user been created
	u, _ := data.GetUserByName(ctx, "user@user")

	if u.UserName == "user@user" {
		log.Println("user@user exist...")
		return
	}
	log.Println("Creating user@user...")
	hashPwd, _ := pkg.HashPassword("password")
	_, err := data.CreateUsers(ctx, gen.CreateUsersParams{
		UserName:     "user@user",
		PassWordHash: hashPwd,
		Name:         "Dummy user",
	})
	if err != nil {
		log.Println("error getting user@dummyuser.domain ", err)
		os.Exit(1)
	}
}
