package main

import (
	"chapter6/internal"
	"chapter6/internal/api"
	"chapter6/store"
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type MockDB store.Queries

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

func (db MockDB) WithTx(tx *sql.Tx) *MockQuery {
	return &MockQuery{}
}

type MockQuery struct {
}

func (db MockDB) GetUserByName(ctx context.Context, userName string) (store.GowebappUser, error) {
	return store.GowebappUser{
		UserID:       0,
		UserName:     "username",
		PasswordHash: "passwordhash",
		Name:         "",
		Config:       nil,
		CreatedAt:    time.Time{},
		IsEnabled:    false,
	}, nil
}

type MockLoginInterface struct{}

func (m MockLoginInterface) GetUserByName(req *http.Request, wr http.ResponseWriter,
	q *store.Queries, payload LoginRequest) {
	api.JSONMessage(wr, http.StatusOK, "")
}

func TestLogin(t *testing.T) {
	mock := MockDB{}
	qq := store.New(mock)
	mockLoginInterface := MockLoginInterface{}

	s := api.NewServer(qq, internal.GetAsInt("SERVER_PORT", 9002))

	s.AddRoute("/login", handleLogin(mockLoginInterface, qq), http.MethodPost)

	str := "{\"username\":\"username\", \"password\":\"password\"}"
	reader := strings.NewReader(str)
	req, _ := http.NewRequest("POST", "/login", reader)
	response := executeRequest(req, s)
	checkResponseCode(t, 200, response.Code)
	checkResponseBody(t, "{\"status\":\"200 / OK\"}\n", response.Body.String())
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseBody(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected body %s. Got %s\n", expected, actual)
	}
}

func executeRequest(req *http.Request, s *api.Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}
