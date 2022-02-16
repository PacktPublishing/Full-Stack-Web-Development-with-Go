package model

import (
	chapter6 "chapter6/gen"
	"context"
	"errors"
)

type Sets struct {
	Weight     int32 `json:"weight"`
	ExerciseId int64 `json:"exerciseid"`
}

type ListSets []Sets

func (s Sets) AddSets(q *chapter6.Queries) error {
	ctx := context.Background()

	ex := Exercise{}
	l := ex.ListExercise(s.ExerciseId, q)

	if l.ExerciseId != 0 {
		_, err := q.UpsertSet(ctx, chapter6.UpsertSetParams{
			ExerciseID: s.ExerciseId,
			Weight:     s.Weight,
		})
		return err
	}

	return errors.New("")
}
