package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// WrapEmptyJSON takes a byte stream and if there's no data
// will allow inserting an empty JSON object. Useful
// when API calls are omitempty
func WrapEmptyJSON(data []byte) []byte {
	if len(data) > 0 {
		return data
	}
	return []byte(EmptyJSON())
}

func EmptyJSON() string {
	return "{}"
}

func JSONError(rw http.ResponseWriter, errorCode int, errorMessages ...string) {
	if len(errorMessages) > 1 {
		rw.WriteHeader(errorCode)
		json.NewEncoder(rw).Encode(struct {
			Status string   `json:"status,omitempty"`
			Errors []string `json:"errors,omitempty"`
		}{
			Status: fmt.Sprintf("%d / %s", errorCode, http.StatusText(errorCode)),
			Errors: errorMessages,
		})
		return
	}

	rw.WriteHeader(errorCode)
	json.NewEncoder(rw).Encode(struct {
		Status string `json:"status,omitempty"`
		Error  string `json:"error,omitempty"`
	}{
		Status: fmt.Sprintf("%d / %s", errorCode, http.StatusText(errorCode)),
		Error:  errorMessages[0],
	})
}

func PrettyJSON(obj interface{}) []byte {
	prettyJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println("Failed to generate json", err)
	}
	return prettyJSON
}
