package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handlerGetHelloWorld(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wr, "Hello, World\n")
	log.Println(req.Method) // request method
	log.Println(req.URL)    // request URL
	log.Println(req.Header) // request headers
	log.Println(req.Body)   // request body)
}

type foo struct {
}

func (f foo) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wr, "Hello, World Alternate\n")
	log.Println(req.Method) // request method
	log.Println(req.URL)    // request URL
	log.Println(req.Header) // request headers
	log.Println(req.Body)   // request body)
}

func main() {
	// Set some flags for easy debugging
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	// Get a port from ENV var or default to 9002
	port := "9002"
	if value, exists := os.LookupEnv("SERVER_PORT"); exists {
		port = value
	}

	// We could use the default mux and we then use http.HandleFunc and http.ListenAndServe(port,nil) but
	// I find it best to create your own. As we'll see with later patterns we will cover graceful
	// shutdown as well.
	router := http.NewServeMux()

	srv := http.Server{
		Addr:           ":" + port, // Addr optionally specifies the listen address for the server in the form of "host:port".
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20, // This is 1MB, it's good practice to limit how much data you'll accept from a client.
	}

	// This is just to show an alternate way to declare a handler
	// by having a struct that implements the ServeHTTP(...) interface
	dummyHandler := foo{}

	router.HandleFunc("/", handlerGetHelloWorld)
	router.Handle("/1", dummyHandler)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Couldnt ListenAndServe()", err)
	}
}
