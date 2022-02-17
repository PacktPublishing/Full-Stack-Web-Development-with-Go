package model

import (
	chapter6 "chapter6/gen"
	"context"
	"errors"
)

type SetsInterface interface {
	AddSets(q *chapter6.Queries) error
	ListSets(q *chapter6.Queries) ListSets
}

type Sets struct {
	Weight     int32 `json:"weight"`
	ExerciseId int64 `json:"exerciseid"`
	SetsId     int64 `json:"setsid"`
}

type ListSets []Sets

func (s Sets) AddSets(q *chapter6.Queries) error {
	ctx := context.Background()

	ex, err := q.ListExercise(ctx, s.ExerciseId)

	if err == nil {
		_, err := q.UpsertSet(ctx, chapter6.UpsertSetParams{
			ExerciseID: ex.ExerciseID,
			Weight:     s.Weight,
		})
		return err
	}

	return errors.New("")
}

func (s Sets) ListSets(q *chapter6.Queries) ListSets {
	ctx := context.Background()
	l, _ := q.ListSets(ctx)

	var lSets ListSets

	for _, e := range l {
		lSets = append(lSets, Sets{
			ExerciseId: e.ExerciseID,
			SetsId:     e.SetID,
			Weight:     e.Weight,
		})
	}

	return lSets
}
