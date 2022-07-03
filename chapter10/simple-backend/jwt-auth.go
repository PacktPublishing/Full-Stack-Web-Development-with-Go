package main

import (
	"encoding/json"
	"errors"
	"go-server/env"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSigningKey []byte
var defaultCookie http.Cookie
var jwtSessionLength time.Duration
var jwtSigningMethod = jwt.SigningMethodHS256

func init() {
	jwtSigningKey = []byte(env.GetAsString("JWT_SIGNING_KEY", "PacktPub"))
	defaultCookie = http.Cookie{
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Domain:   env.GetAsString("COOKIE_DOMAIN", "localhost"),
		Secure:   env.GetAsBool("COOKIE_SECURE", true),
	}
	jwtSessionLength = time.Duration(env.GetAsInt("JWT_SESSION_LENGTH", 5))
}

// JWTProtectedMiddleware verifies a valid JWT exists in our
// cookie and if not, encourages the consumer to login again.
func JWTProtectedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Grab jwt-token cookie
		jwtCookie, err := r.Cookie("jwt-token")
		if err != nil {
			log.Println("Error occurred reading cookie", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message,omitempty"`
			}{
				Message: "Your session is not valid - please login",
			})
			return
		}

		// You can examine the raw cookie here if you uncomment the
		// following line.
		// log.Println("Got cookie value", jwtCookie.Value)

		// Decode and validate JWT if there is one
		userEmail, err := decodeJWTToUser(jwtCookie.Value)

		if userEmail == "" || err != nil {
			log.Println("Error decoding token", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message,omitempty"`
			}{
				Message: "Your session is not valid - please login",
			})
			return
		}

		// If it's good, update the expiry time
		freshToken := createJWTTokenForUser(userEmail)

		// Set the new cookie and continue into the handler
		w.Header().Add("Content-Type", "application/json")
		http.SetCookie(w, authCookie(freshToken))
		next.ServeHTTP(w, r)
	})
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func createJWTTokenForUser(user string) string {
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * jwtSessionLength).Unix(),
		},
	}

	// Encode to token string
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSigningKey)
	if err != nil {
		log.Println("Error occurred generating JWT", err)
		return ""
	}
	return tokenString
}

func decodeJWTToUser(token string) (string, error) {
	// Decode
	decodeToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if !(jwtSigningMethod == token.Method) {
			// Check our method hasn't changed since issuance
			return nil, errors.New("signing method mismatch")
		}
		return jwtSigningKey, nil
	})

	// There's two parts. We might decode it successfully but it might
	// be the case we aren't Valid so you must check both
	if decClaims, ok := decodeToken.Claims.(*CustomClaims); ok && decodeToken.Valid {
		return decClaims.Username, nil
	}
	return "", err
}

func authCookie(token string) *http.Cookie {
	d := defaultCookie
	d.Name = "jwt-token"
	d.Value = token
	d.Path = "/"
	return &d
}

func expiredAuthCookie() *http.Cookie {
	d := defaultCookie
	d.Name = "jwt-token"
	d.Value = ""
	d.Path = "/"
	d.MaxAge = -1
	d.Expires = time.Date(1983, 7, 26, 20, 34, 58, 651387237, time.UTC)
	return &d
}
