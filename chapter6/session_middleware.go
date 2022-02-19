package main

import (
	"chapter6/internal/api"
	"chapter6/store"
	"context"
	"database/sql"
	"errors"
	"net/http"
)

// This middleware isn't reusable outside of our app as it uses
// generated stores and is best extracted and kept coupled to the rest
// of our code

type UserSession struct {
	UserID int64
}

// We define this so it can't clash outside our package
// with anything else.
type ourCustomKey string

const SessionKey ourCustomKey = "unique-session-key-for-our-example"

// Our custom middleware to ensure
// we have a valid user session
func validCookieMiddleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
			session, err := cookieStore.Get(req, "session-name")
			if err != nil {
				api.JSONError(wr, http.StatusInternalServerError, "Session Error")
				return
			}

			userID := session.Values["userID"].(int64)
			isAuthd := session.Values["userAuthenticated"].(bool)

			if !isAuthd || userID < 1 {
				api.JSONError(wr, http.StatusForbidden, "Bad Credentials")
				return
			}

			querier := store.New(db)
			user, err := querier.GetUser(req.Context(), int64(userID))
			if errors.Is(err, sql.ErrNoRows) {
				api.JSONError(wr, http.StatusForbidden, "Bad Credentials")
				return
			}

			ctx := context.WithValue(req.Context(), SessionKey, UserSession{
				UserID: user.UserID,
			})
			h.ServeHTTP(wr, req.WithContext(ctx))
		})
	}
}

func userFromSession(req *http.Request) (UserSession, bool) {
	session, ok := req.Context().Value(SessionKey).(UserSession)
	if session.UserID < 1 {
		// Shouldnt happen
		return UserSession{}, false
	}
	return session, ok
}
