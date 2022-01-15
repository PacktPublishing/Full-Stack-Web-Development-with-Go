package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handlerSlug(wr http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)["slug"]
	if slug == "" {
		log.Println("Slug not provided")
		return
	}
	log.Println("Got slug", slug)
}

func handlerGetHelloWorld(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wr, "Hello, World\n")
	log.Println("Request via", req.Method) // request method
	log.Println(req.URL)                   // request URL
	log.Println(req.Header)                // request headers
	log.Println(req.Body)                  // request body)
}

func handlerPostEcho(wr http.ResponseWriter, req *http.Request) {
	log.Println("Request via", req.Method) // request method
	log.Println(req.URL)                   // request URL
	log.Println(req.Header)                // request headers

	// We are going to read it into a buffer
	// as the request body is an io.ReadCloser
	// and so we should only read it once.
	body, err := ioutil.ReadAll(req.Body)

	log.Println("read >", string(body), "<")

	n, err := io.Copy(wr, bytes.NewReader(body))
	if err != nil {
		log.Println("Error echoing response", err)
	}
	log.Println("Wrote back", n, "bytes")
}

func main() {
	// Set some flags for easy debugging
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	// Get a port from ENV var or default to 9002
	port := "9002"
	if value, exists := os.LookupEnv("SERVER_PORT"); exists {
		port = value
	}

	// Off the bat, we can enforce StrictSlash
	// This is a nice helper function that means
	// When true, if the route path is "/foo/", accessing "/foo" will perform a 301 redirect to the former and vice versa.
	// In other words, your application will always see the path as specified in the route.
	// When false, if the route path is "/foo", accessing "/foo/" will not match this route and vice versa.

	router := mux.NewRouter().StrictSlash(true)

	srv := http.Server{
		Addr:    ":" + port, // Addr optionally specifies the listen address for the server in the form of "host:port".
		Handler: router,
	}

	router.HandleFunc("/", handlerGetHelloWorld).Methods(http.MethodGet)
	router.HandleFunc("/", handlerPostEcho).Methods(http.MethodPost)
	router.HandleFunc("/{slug}", handlerSlug).Methods(http.MethodGet)

	log.Println("Starting on", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Couldnt ListenAndServe()", err)
	}
}
