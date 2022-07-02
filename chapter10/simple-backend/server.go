package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := ":8000"

	rtr := mux.NewRouter()
	rtr.Handle("/", appGET()).Methods(http.MethodGet)
	rtr.Handle("/login", appPOSTLogin()).Methods(http.MethodPost, http.MethodOptions)
	rtr.HandleFunc("/logout", func(rw http.ResponseWriter, req *http.Request) {
		http.SetCookie(rw, expiredAuthCookie())
	}).Methods(http.MethodGet)

	rtr.Use(
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type"}),
			handlers.AllowedOrigins([]string{"http://0.0.0.0:3000", "http://localhost:3000"}),
			handlers.AllowCredentials(),
			handlers.AllowedMethods([]string{
				http.MethodGet,
				http.MethodPost,
				http.MethodOptions,
			})),
	)

	// Our private router
	privateRtr := rtr.PathPrefix("/private").Subrouter()
	privateRtr.Handle("/", appGETPrivate()).Methods(http.MethodGet)
	privateRtr.Handle("/", appPOSTPrivate()).Methods(http.MethodPost, http.MethodOptions)
	privateRtr.Use(
		JWTProtectedMiddleware,
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type"}),
			handlers.AllowedOrigins([]string{"http://0.0.0.0:3000", "http://localhost:3000"}),
			handlers.AllowCredentials(),
			handlers.AllowedMethods([]string{
				http.MethodGet,
				http.MethodPost,
				http.MethodOptions,
			})),
	)

	// Apply the CORS middleware to our top-level router, with the defaults.
	log.Printf("Listening on http://0.0.0.0%s/", port)
	http.ListenAndServe(port, rtr)
}

func appGET() http.HandlerFunc {
	type ResponseBody struct {
		Response string `json:"message,omitempty"`
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		log.Println("GET", req)

		json.NewEncoder(rw).Encode(ResponseBody{
			Response: "Hello World",
		})
	}
}

func appGETPrivate() http.HandlerFunc {
	type ResponseBody struct {
		Response string `json:"message,omitempty"`
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		log.Println("GET", req)

		json.NewEncoder(rw).Encode(ResponseBody{
			Response: "Hello World from /private",
		})
	}
}

func appPOSTPrivate() http.HandlerFunc {
	type RequestBody struct {
		Message string `json:"message,omitempty"`
	}
	type ResponseBody struct {
		Response string `json:"response,omitempty"`
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		log.Println("POST", req)

		var rb RequestBody
		if err := json.NewDecoder(req.Body).Decode(&rb); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("We received an inbound value of", rb.Message)

		json.NewEncoder(rw).Encode(ResponseBody{
			Response: "ok",
		})
	}
}

func appPOSTLogin() http.HandlerFunc {
	type RequestBody struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}
	type ResponseBody struct {
		Response string `json:"response,omitempty"`
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		log.Println("POST", req)

		var rb RequestBody
		if err := json.NewDecoder(req.Body).Decode(&rb); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("We received an inbound value of", rb.Username, rb.Password)

		if rb.Username == "admin" && rb.Password == "password" {
			freshToken := createJWTTokenForUser(rb.Username)
			http.SetCookie(rw, authCookie(freshToken))

			json.NewEncoder(rw).Encode(ResponseBody{
				Response: "success",
			})
			return
		}

		json.NewEncoder(rw).Encode(ResponseBody{
			Response: "incorrect username/password",
		})
	}
}
