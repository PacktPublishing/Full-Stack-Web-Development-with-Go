package api

import (
	"mime"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
)

// JSON middleware will help us only handle JSON
// in and out
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")

		if strings.TrimSpace(contentType) == "" {
			var parseError error
			contentType, _, parseError = mime.ParseMediaType(contentType)
			if parseError != nil {
				JSONError(wr, http.StatusBadRequest, "Bad or no content-type header found")
				return
			}
		}

		if contentType != "application/json" {
			JSONError(wr, http.StatusUnsupportedMediaType, "Content-Type not application/json")
			return
		}
		// Tell the client we're talking JSON as well.
		wr.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(wr, req)
	})
}

func CORSMiddleware(origins []string) func(http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedHeaders([]string{
			"X-Requested-With", "Authorization", "Access-Control-Allow-Methods",
			"Access-Control-Allow-Origin", "Origin", "Accept", "Content-Type",
		}),
		handlers.AllowedOrigins(origins),
		handlers.AllowedMethods([]string{
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
		}),
	)
}
