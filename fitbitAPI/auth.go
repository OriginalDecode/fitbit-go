package fitbitAPI

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type AuthenticationData struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

var oauth2Config *oauth2.Config

func login(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("auth")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func getAuthKey() string {
	return base64.StdEncoding.EncodeToString([]byte(oauth2Config.ClientID + ":" + oauth2Config.ClientSecret))
}

func callback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Callback")
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read response body:", err.Error())
		log.Println(body)
	}

	var data AuthenticationData
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Println("Failed to unmarshal initial auth data:", err.Error())
	}

	requestBody := url.Values{}
	requestBody.Set("code", data.Code)
	requestBody.Set("grant_type", "authorization_code")
	requestBody.Set("client_id", oauth2Config.ClientID)
	requestBody.Set("redirect_uri", oauth2Config.RedirectURL)

	req, err := http.NewRequest("POST", oauth2Config.Endpoint.TokenURL, strings.NewReader(requestBody.Encode()))
	if err != nil {
		log.Println("Failed to get token:", err.Error())
	}

	key := getAuthKey()

	req.Header.Set("Authorization", "Basic "+key)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var httpClient = http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Failed to authorize using basic key:", err.Error())
	}
	defer resp.Body.Close()
	fmt.Println("Request Sent:")

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read body", err.Error())
	}

	authResp, err := UnmarshalAuth([]byte(body))
	if err != nil {
		log.Println("Failed to unmarshal auth data:", err.Error())
	} else {
		fmt.Println("Success")
	}
	fmt.Println("{}", authResp)

	// http.Redirect(w, r, "http://127.0.0.1:8081/access", 200)

	// var htmlIndex = `<html>
	// <body>
	// 	<p>logged in</p>
	// 	<a href="/request">Request Data</a>
	// </body>
	// </html>`
	// fmt.Fprintf(w, htmlIndex)
}

func Auth(config *oauth2.Config, urlPath string, callbackUrl string) {
	oauth2Config = config
	http.HandleFunc(urlPath, login)
	http.HandleFunc(callbackUrl, callback)
}
