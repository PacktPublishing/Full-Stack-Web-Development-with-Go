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
	return []byte("{}}")
}

func JSONError(wr http.ResponseWriter, errorCode int, errorMessages ...string) {
	wr.WriteHeader(errorCode)
	if len(errorMessages) > 1 {
		json.NewEncoder(wr).Encode(struct {
			Status string   `json:"status,omitempty"`
			Errors []string `json:"errors,omitempty"`
		}{
			Status: fmt.Sprintf("%d / %s", errorCode, http.StatusText(errorCode)),
			Errors: errorMessages,
		})
		return
	}

	json.NewEncoder(wr).Encode(struct {
		Status string `json:"status,omitempty"`
		Error  string `json:"error,omitempty"`
	}{
		Status: fmt.Sprintf("%d / %s", errorCode, http.StatusText(errorCode)),
		Error:  errorMessages[0],
	})
}

func JSONMessage(wr http.ResponseWriter, code int, messages ...string) {
	wr.WriteHeader(code)
	if len(messages) > 1 {
		json.NewEncoder(wr).Encode(struct {
			Status   string   `json:"status,omitempty"`
			Messages []string `json:"messages,omitempty"`
		}{
			Status:   fmt.Sprintf("%d / %s", code, http.StatusText(code)),
			Messages: messages,
		})
		return
	}

	json.NewEncoder(wr).Encode(struct {
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	}{
		Status:  fmt.Sprintf("%d / %s", code, http.StatusText(code)),
		Message: messages[0],
	})
}

func PrettyJSON(obj interface{}) []byte {
	prettyJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println("Failed to generate json", err)
	}
	return prettyJSON
}
