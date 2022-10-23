package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/horitaka/oidc-cli/constants"
	"github.com/joho/godotenv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: リファクタリング
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load env file: %v", err)
	}

	u, err := url.Parse(constants.AUTH_URL)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	clienId := os.Getenv("CLIENT_ID")
	q.Set("client_id", clienId)
	q.Set("scope", "openid email profile")
	q.Set("response_type", "code")
	q.Set("redirect_uri", constants.CALLBACK_URL)
	q.Set("state", "state") // TODO
	q.Set("nonce", "nonce") // TODO

	u.RawQuery = q.Encode()
	// fmt.Println(u)

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("location", u.String())
	w.WriteHeader(http.StatusFound)

	// TODO: login後にプログラムを終了する
	// os.Exit(1)
}
