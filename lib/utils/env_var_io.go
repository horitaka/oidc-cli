package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type AuthClient struct {
	ClientId     string
	ClientSecret string
}

func LoadEnv() AuthClient {
	// fmt.Println(os.Getenv("CLIENT_ID"))

	_, err := os.Stat(".env")
	if err == nil {
		godotenv.Load(".env")
	}

	authClient := AuthClient{
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
	return authClient
}
