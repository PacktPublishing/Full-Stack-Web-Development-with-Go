package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

var serviceA, serviceB bool

func main() {
	port := ":8000"
	rtr := mux.NewRouter()
	rtr.Handle("/", handler()).Methods(http.MethodGet)

	log.Printf("Listening on http://0.0.0.0%s/", port)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(w *sync.WaitGroup) {
		defer w.Done()
		serviceA = checkFlags("servicea")
		serviceB = checkFlags("serviceb")
	}(wg)
	wg.Wait()

	http.ListenAndServe(port, rtr)
}

func checkFlags(key string) bool {
	type FeatureFlagServerResponse struct {
		Enabled bool `json:"enabled"`
	}
	requestURL := fmt.Sprintf("http://localhost:%d/features/%s", 8080, key)
	res, err := http.Get(requestURL)
	if err != nil {
		log.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var f FeatureFlagServerResponse
	err = json.Unmarshal(resBody, &f)
	if err != nil {
		log.Printf("Error unmarshalling: %s\n", err)
		os.Exit(1)
	}

	return f.Enabled
}

func handler() http.HandlerFunc {
	type ResponseBody struct {
		Message string
	}
	return func(rw http.ResponseWriter, req *http.Request) {
		var a, b string
		if serviceA {
			a = callService("8081")
		}
		if serviceB {
			b = callService("8082")
		}

		json.NewEncoder(rw).Encode(ResponseBody{
			Message: a + "-" + b,
		})
	}
}

func callService(port string) string {
	type ServiceResponse struct {
		Message string
	}
	requestURL := fmt.Sprintf("http://localhost:%s", port)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	var f ServiceResponse
	err = json.Unmarshal(resBody, &f)
	if err != nil {
		log.Printf("Error unmarshalling: %s\n", err)
		os.Exit(1)
	}

	return f.Message
}
