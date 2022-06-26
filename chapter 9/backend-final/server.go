package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/example/stringutil"
)

func main() {
	port := ":8000"

	rtr := mux.NewRouter()
	rtr.Handle("/", appGET()).Methods(http.MethodGet)
	rtr.Handle("/", appPOST()).Methods(http.MethodPost, http.MethodOptions)

	rtr.Use(
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Origin", "Content-Type"}),
			handlers.AllowedOrigins([]string{"http://0.0.0.0:3000", "http://localhost:3000"}),
			handlers.AllowCredentials(),
			handlers.AllowedMethods([]string{
				http.MethodGet,
				http.MethodPost,
			})),
	)

	// Apply the CORS middleware to our top-level router, with the defaults.
	log.Printf("Listening on http://0.0.0.0%s/", port)
	http.ListenAndServe(port, rtr)
}

func appGET() http.HandlerFunc {
	type ResponseBody struct {
		Message string `json:"message,omitempty"`
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		log.Println("GET", req)
		json.NewEncoder(rw).Encode(ResponseBody{
			Message: "Hello World",
		})
	}
}

func appPOST() http.HandlerFunc {
	type RequestBody struct {
		Inbound string `json:"inbound,omitempty"`
	}
	type ResponseBody struct {
		Outbound string `json:"outbound,omitempty"`
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		log.Println("POST", req)

		var rb RequestBody
		if err := json.NewDecoder(req.Body).Decode(&rb); err != nil {
			log.Println("apiAdminPatchUser: Decode failed:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("We received an inbound value of", rb.Inbound)
		json.NewEncoder(rw).Encode(ResponseBody{
			Outbound: stringutil.Reverse(rb.Inbound),
		})
	}
}
