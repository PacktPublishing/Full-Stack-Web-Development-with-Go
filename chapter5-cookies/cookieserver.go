package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func setCookie() http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.PostFormValue("username") == "user@user" && req.PostForm.Get("password") == "password" {
			log.Println("Setting cookie")
			// Note: Set the cookie before writing the response
			http.SetCookie(wr, &http.Cookie{
				Name:  "user-session",
				Value: "user@user:password",
			})
			fmt.Fprintf(wr, "Successful login")
			return
		}

		fmt.Fprintf(wr, "Bad login")
	}
}
func checkCookie() http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		log.Println("Checking cookies:")
		for _, c := range req.Cookies() {
			log.Println(c)
		}
	}
}
func unsetCookie() http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		log.Println("Deleting cookie")
		http.SetCookie(wr, &http.Cookie{
			Name:   "user-session",
			Value:  "",
			MaxAge: -1,
			Expires: time.Date(
				1983, 7, 26, 20, 34, 58, 651387237, time.UTC),
		})
		fmt.Fprintf(wr, "Successful logout")
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/login", setCookie()).Methods(http.MethodPost)
	router.HandleFunc("/", checkCookie()).Methods(http.MethodGet)
	router.HandleFunc("/logout", unsetCookie()).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("starting server on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
