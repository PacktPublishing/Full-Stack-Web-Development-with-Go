package main

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func main() {
	mySigningKey := []byte("PacktPub")

	// Your claims above and beyond the default
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		// Note we embed the standard claims here
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "FullStackGo",
		},
	}

	// Encode to token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	log.Printf("Your JWT as a string is %v - %v", tokenString, err)

	// Try these different strings - you can generate more at https://jwt.io
	// tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJkYXZlIiwiZXhwIjoxMjN9.V_X0h79nM_BwlJkACEZQn_WH7kH6CIJEh1zsv02caQM"
	// tokenString = "eyJhbGciOiJIUzUxMiJ9.e30.bjEexYLMl8_5CYqvJRSztlZHPBRonj5bXw4c0FwsoWtcGL1olxYjYqMxJZLjPKH5Yq6kboCABaypgpOxFfPAgg"
	// tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjEwNTc4OTQwNjAsImlzcyI6InRlc3QifQ.WSCRl6od_Wi4h5f8k5-vGxZZPWkDafvP2rDtAR8gdw8"
	// tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE2NTY3MzU1NjN9.hboaRqGbOaUbqMoaD4wRlNxq7CIMAGJ62t2KOYoKW3k"

	// Decode
	decodeToken, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		ok := jwt.SigningMethodHS256 == token.Method
		if !ok {
			return nil, errors.New("Signing Method mismatch!")
		}
		return mySigningKey, nil
	})

	// There's two parts. We might decode it successfully but it might
	// be the case we aren't Valid
	if decClaims, ok := decodeToken.Claims.(*MyCustomClaims); ok && decodeToken.Valid {
		log.Println("Issuer match?", decClaims.VerifyIssuer("FullStackGo", true))
		log.Printf("%v %v", decClaims.Foo, decClaims.StandardClaims.ExpiresAt)
	} else {
		log.Println("Error", err)
	}
	log.Println("finished")
}
