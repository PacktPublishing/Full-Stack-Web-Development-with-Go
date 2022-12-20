package main

import (
	"context"
	"database/sql"
	chapter1 "fitness.dev/app/gen"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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

	_, err = st.CreateUsers(ctx, chapter1.CreateUsersParams{
		UserName:     "testuser",
		PassWordHash: "hash",
		Name:         "test",
	})

	if err != nil {
		log.Fatalln("Error creating user :", err)
	}

	eid, err := st.CreateExercise(ctx, "Exercise1")

	if err != nil {
		log.Fatalln("Error creating exercise :", err)
	}

	set, err := st.CreateSet(ctx, chapter1.CreateSetParams{
		ExerciseID: eid,
		Weight:     100,
	})

	if err != nil {
		log.Fatalln("Error updating exercise :", err)
	}

	set, err = st.UpdateSet(ctx, chapter1.UpdateSetParams{
		ExerciseID: eid,
		SetID:      set.SetID,
		Weight:     2000,
	})

	if err != nil {
		log.Fatalln("Error updating set :", err)
	}

	log.Println("Done!")

	u, err := st.ListUsers(ctx)

	for _, usr := range u {
		fmt.Println(fmt.Sprintf("Name : %s, ID : %d", usr.Name, usr.UserID))
	}
}
