package main

import (
	"./fitbitAPI"
	// "encoding/base64"
	// "encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/fitbit"
	// "io/ioutil"
	"log"
	"net/http"
	// "net/url"
	"os"
	// "strings"
	// "time"
)

const (
	API_URL string = "https://api.fitbit.com"
)

var oauth2Config *oauth2.Config

func main() {
	fmt.Println("Starting fitbit-fetcher")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	scopes := fitbitAPI.GetScopes()
	oauth2Config = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:8080/callback/",
		ClientID:     os.Getenv("FITBITAPI_CLIENT_ID"),
		ClientSecret: os.Getenv("FITBITAPI_CLIENT_SECRET"),
		Endpoint:     fitbit.Endpoint,
		Scopes:       []string{scopes.Profile, scopes.Weight},
	}

	fitbitAPI.Auth(oauth2Config, "/login", "/callback/")

	http.HandleFunc("/v1/request", fitbitAPI.Request)
	http.ListenAndServe(":"+os.Getenv("FITBITAPI_PORT"), nil)

}
