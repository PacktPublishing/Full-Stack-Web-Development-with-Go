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
			handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Origin", "Accept", "Content-Type", "business-id", "X-Device"}),
			handlers.AllowedOrigins([]string{"http://0.0.0.0:3000", "http://localhost:3000"}),
			handlers.AllowCredentials(),
			handlers.AllowedMethods([]string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
				http.MethodOptions,
				http.MethodPatch,
			})),
	)

	// Apply the CORS middleware to our top-level router, with the defaults.
	log.Printf("Listening on http://0.0.0.0%s/", port)
	http.ListenAndServe(port, rtr)
}

func appGET() http.HandlerFunc {
	type ResponseBody struct {
		Message string
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
		Inbound string
	}
	type ResponseBody struct {
		OutBound string
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
			OutBound: stringutil.Reverse(rb.Inbound),
		})
	}
}
