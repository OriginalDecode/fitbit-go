package main

import (
	"./fitbitAPI"
	// "encoding/base64"
	// "encoding/json"
	"fmt"
	// "github.com/gorilla/schema"
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
	ScopeActivity  = "activity"
	ScopeHeartrate = "heartrate"
	ScopeLocation  = "location"
	ScopeNutrition = "nutrition"
	ScopeProfile   = "profile" // is this what I request?
	ScopeSettings  = "settings"
	ScopeSleep     = "sleep"
	ScopeSocial    = "social"
	ScopeWeight    = "weight"
)

const (
	API_URL string = "https://api.fitbit.com"
)

var AuthData fitbitAPI.AuthenticationData
var RespData fitbitAPI.AuthResponse
var Client = http.DefaultClient
var oauth2Config *oauth2.Config

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
<body>
	<a href="/login">Fitbit Log In</a>
</body>
</html>`
	fmt.Fprintf(w, htmlIndex)
}

func main() {
	fmt.Println("Starting fitbit-fetcher")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	oauth2Config = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:8080/callback/",
		ClientID:     os.Getenv("FITBITAPI_CLIENT_ID"),
		ClientSecret: os.Getenv("FITBITAPI_CLIENT_SECRET"),
		Endpoint:     fitbit.Endpoint,
		Scopes:       []string{ScopeProfile},
	}

	fitbitAPI.Auth(oauth2Config, "/login", "/callback/")

	http.HandleFunc("/", handleMain)
	// http.HandleFunc("/callback/", handleCallback)
	http.ListenAndServe(":8080", nil)

}
