package main

import (
	"database/sql"
	"net/http"
)

func handleLogin(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {

	})
}

func handleLogout() http.HandlerFunc {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		// TODO - delete cookie
	})
}
