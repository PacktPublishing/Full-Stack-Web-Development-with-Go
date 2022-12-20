package main

import (
	"bytes"
	"encoding/json"
	"log"
)

// example to show logging using standard library
func main() {
	ol := log.Default()

	// set log format to - dd/mm/yy hh:mm:ss
	ol.SetFlags(log.LstdFlags)
	ol.Println("Just a log text")
	lognumber(ol)
	logjson(ol)
}

// logjson to log json to logger
func logjson(ol *log.Logger) {
	ol.SetFlags(log.Ltime)

	ex := `{"name": "Cake","batters":{"batter":[{ "id": "001", "type": "Good Food" }]},"topping":[{ "id": "002", "type": "Syrup" }]}`

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(ex), "", "\t")
	if error != nil {
		ol.Fatalf("Error parsing : %s", error.Error())
	}

	ol.Println(string(prettyJSON.Bytes()))
}

// lognumber to log number to logger
func lognumber(ol *log.Logger) {
	ol.SetFlags(log.Lshortfile) // will display filename:linenumber format
	ol.Printf("This is number %d", 1)
}
