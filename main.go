package main

import (
	"./fitbitAPI"
	"encoding/base64"
	"encoding/json"
	"fmt"
	// "github.com/gorilla/schema"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/fitbit"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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
	API_URL   string = "https://api.fitbit.com"
	TOKEN_URL string = "https://api.fitbit.com/oauth2/token"
	AUTH_URL  string = "https://www.fitbit.com/oauth2/authorize"
)

type AuthenticationData struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

var AuthData AuthenticationData
var RespData fitbitAPI.AuthResponse
var CLIENT = http.DefaultClient
var oauth2Config *oauth2.Config

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
<body>
	<a href="/get/login">Fitbit Log In</a>
</body>
</html>`
	fmt.Fprintf(w, htmlIndex)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("auth")
	log.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {

	log.Print("handleCallback")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("{}", err.Error())
	}

	var data AuthenticationData
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("{}", err.Error())
	}

	requestBody := url.Values{}
	requestBody.Set("code", data.Code)
	requestBody.Set("grant_type", "authorization_code")
	requestBody.Set("client_id", oauth2Config.ClientID)
	requestBody.Set("redirect_uri", "http://127.0.0.1:8081/callback")

	req, err := http.NewRequest("POST", TOKEN_URL, strings.NewReader(requestBody.Encode()))
	if err != nil {
		panic(err.Error())
	}

	key := base64.StdEncoding.EncodeToString([]byte(oauth2Config.ClientID + ":" + oauth2Config.ClientSecret))

	req.Header.Set("Authorization", "Basic "+key)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := CLIENT.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	fmt.Println("Request Sent...")

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	authResp, err := fitbitAPI.UnmarshalAuth([]byte(body))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(authResp)

	http.Redirect(w, r, "http://127.0.0.1:8081/access", 200)

	var htmlIndex = `<html>
	<body>
		<p>logged in</p>
		<a href="/request">Request Data</a>
	</body>
	</html>`
	fmt.Fprintf(w, htmlIndex)
}

func main() {
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

	fmt.Println("Starting fitbit-fetcher")
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/get/login", handleLogin)
	http.HandleFunc("/callback/", handleCallback)
	http.ListenAndServe(":8080", nil)

}
