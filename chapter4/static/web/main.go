package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Starting up server on port 3333 ...")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal("error occurred starting up server : ", err)
	}
}
