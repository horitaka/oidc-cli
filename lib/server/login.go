package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

const AUTH_BASE_URL = "https://accounts.google.com/o/oauth2/v2/auth"

func Login(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load env file: %v", err)
	}

	u, err := url.Parse(AUTH_BASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	clienId := os.Getenv("CLIENT_ID")
	q.Set("client_id", clienId)
	q.Set("scope", "openid")
	q.Set("response_type", "code")
	q.Set("redirect_uri", "http://localhost:8080/callback")
	q.Set("state", "state") // TODO
	q.Set("nonce", "nonce") // TODO

	u.RawQuery = q.Encode()
	// fmt.Println(u)

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", u.String())
	w.WriteHeader(http.StatusFound)
}
