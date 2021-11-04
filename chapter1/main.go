package main

import (
	"context"
	"database/sql"
	chapter1 "fitness.dev/app/gen"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {
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
	st := chapter1.New(db)

	ctx := context.Background()

	chuser, err := st.CreateUsers(ctx, chapter1.CreateUsersParams{
		UserName:     "testuser",
		PassWordHash: "hash",
		Name:         "test",
	})

	if (err!=nil) {
		os.Exit(1)
	}

	eid, err := st.CreateExercise(ctx, "Exercise1")

	if (err!=nil) {
		os.Exit(1)
	}

	sid, err := st.UpsertSet(ctx, chapter1.UpsertSetParams{
		ExerciseID: eid,
		Weight:     100,
	})

	if (err!=nil) {
		os.Exit(1)
	}
	
	st.UpsertWorkout(ctx, chapter1.UpsertWorkoutParams{
		UserID:    chuser.UserID,
		SetID:     sid,
		StartDate: time.Time{},
	})
}
