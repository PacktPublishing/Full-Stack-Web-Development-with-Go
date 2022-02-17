package main

import (
	chapter6 "chapter6/gen"
	"chapter6/model"
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDB struct{}

func (db MockDB) ExecContext(context.Context, string, ...interface{}) (sql.Result, error) {
	return nil, nil
}
func (db MockDB) PrepareContext(context.Context, string) (*sql.Stmt, error) {
	return nil, nil

}
func (db MockDB) QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error) {
	return nil, nil

}
func (db MockDB) QueryRowContext(context.Context, string, ...interface{}) *sql.Row {
	return nil
}

type MockSets struct{}

func (s MockSets) AddSets(q *chapter6.Queries) error {
	return nil
}

func (s MockSets) ListSets(q *chapter6.Queries) model.ListSets {
	return nil
}

func TestListSets(t *testing.T) {
	s := NewServer(chapter6.New(MockDB{}))
	s.hListSetsFn = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
	s.SetupRoutes()
	req, _ := http.NewRequest("GET", "/listsets", nil)
	response := executeRequest(req, s)
	checkResponseCode(t, 200, response.Code)
}

func TestLogin(t *testing.T) {
	s := NewServer(chapter6.New(MockDB{}))
	s.hLogin = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
	s.SetupRoutes()
	req, _ := http.NewRequest("POST", "/login", nil)
	response := executeRequest(req, s)
	checkResponseCode(t, 200, response.Code)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request, s server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.router.ServeHTTP(rr, req)

	return rr
}
