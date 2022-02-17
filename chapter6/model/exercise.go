package model

import (
	"chapter6/gen"
	"context"
)

type ExerciseInterface interface {
	AddExercise(q *chapter6.Queries) error
	ListExercise(i int64, q *chapter6.Queries) Exercise
	ListExercises(q *chapter6.Queries) ListExercise
}

type Exercise struct {
	ExerciseId   int64  `json:"exerciseid"`
	ExerciseName string `json:"exercise"`
}

type ExerciseResponse Response

type ListExercise []Exercise

func (e Exercise) AddExercise(q *chapter6.Queries) error {
	ctx := context.Background()
	_, err := q.UpsertExercise(ctx, e.ExerciseName)

	return err
}

func (e Exercise) ListExercise(i int64, q *chapter6.Queries) Exercise {
	ctx := context.Background()
	l, _ := q.ListExercise(ctx, i)

	return Exercise{
		ExerciseId:   l.ExerciseID,
		ExerciseName: l.ExerciseName,
	}
}

func (e Exercise) ListExercises(q *chapter6.Queries) ListExercise {
	ctx := context.Background()
	l, _ := q.ListExercises(ctx)

	var lexercise ListExercise

	for _, ex := range l {
		lexercise = append(lexercise, Exercise{
			ExerciseName: ex.ExerciseName,
			ExerciseId:   ex.ExerciseID,
		})
	}

	return lexercise
}
