package fitbitAPI

import (
	"encoding/base64"
	"encoding/json"
	// "github.com/gorilla/schema"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
)

type ResponseError struct {
	ErrorType string `json:"errorType"`
	Message   string `json:"message"`
}

type AuthResponse struct {
	AccessToken  string          `json:"access_token"`
	ExpiresIn    int32           `json:"expires_in"`
	RefreshToken string          `json:"refresh_token"`
	Scopes       string          `json:"scope"`
	TokenType    string          `json:"token_type"`
	UserId       string          `json:"user_id"`
	Success      bool            `json:"success"`
	Errors       []ResponseError `json:"errors"`
}

var oauth2Config *oauth2.Config
var authData AuthResponse

func GetAccessToken() string {
	return authData.AccessToken
}

func login(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("auth")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func getAuthKey() string {
	return base64.StdEncoding.EncodeToString([]byte(oauth2Config.ClientID + ":" + oauth2Config.ClientSecret))
}

func prepareRequest(authCode string) *http.Request {
	if len(authCode) <= 0 {
		log.Println("No valid auth code was provided", authCode)
		os.Exit(1)
	}
	log.Println("AuthCode:", authCode)
	requestBody := url.Values{}
	requestBody.Set("code", authCode)
	requestBody.Set("grant_type", "authorization_code")
	requestBody.Set("client_id", oauth2Config.ClientID)
	requestBody.Set("redirect_uri", oauth2Config.RedirectURL)
	req, err := http.NewRequest("POST", oauth2Config.Endpoint.TokenURL,
		strings.NewReader(requestBody.Encode()))
	if err != nil {
		log.Println("Failed to get token:", err.Error())
	}

	req.Header.Set("Authorization", "Basic "+getAuthKey())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}

func authRequest(request *http.Request) []byte {

	var httpClient = http.DefaultClient
	resp, err := httpClient.Do(request)
	if err != nil {
		log.Println("Failed to send auth request:", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read body", err.Error())
	}

	return body
}

func callback(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	request := prepareRequest(query.Get("code"))
	body := authRequest(request)

	data, err := unmarshalAuth([]byte(body))
	authData = data

	s := reflect.ValueOf(&authData).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	if err != nil {
		log.Println("Failed to unmarshal auth data:", err.Error())
	}

	http.Redirect(w, r, "/user", 302)
}

func unmarshalAuth(data []byte) (AuthResponse, error) {
	var auth AuthResponse
	err := json.Unmarshal(data, &auth)
	return auth, err
}

/*
 */
func Auth(config *oauth2.Config, urlPath string, callbackUrl string) {
	oauth2Config = config
	http.HandleFunc(urlPath, login)
	http.HandleFunc(callbackUrl, callback)
}
