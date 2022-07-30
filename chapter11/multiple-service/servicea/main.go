package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := ":8081"
	rtr := mux.NewRouter()
	rtr.Handle("/", appGET()).Methods(http.MethodGet)

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
			Message: "ServiceA active",
		})
	}
}
