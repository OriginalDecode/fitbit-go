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

/*
 */
func login(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("auth")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

/*
 */
func getAuthKey() string {
	return base64.StdEncoding.EncodeToString([]byte(oauth2Config.ClientID + ":" + oauth2Config.ClientSecret))
}

func parseCallback(r *http.Request) AuthenticationData {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read response body:", err.Error())
	}
	var data AuthenticationData
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Println("Failed to unmarshal initial auth data:", err.Error())
	}
	return data
}

/*
	returns the body of the response
*/
func authRequest(authCode string) []byte {
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

	var httpClient = http.DefaultClient
	resp, err := httpClient.Do(req)
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

/*
 */
func callback(w http.ResponseWriter, r *http.Request) {
	data := parseCallback(r)
	body := authRequest(data.Code)

	authResp, err := UnmarshalAuth([]byte(body))
	if err != nil {
		log.Println("Failed to unmarshal auth data:", err.Error())
	}
}

/*
 */
func Auth(config *oauth2.Config, urlPath string, callbackUrl string) {
	oauth2Config = config
	http.HandleFunc(urlPath, login)
	http.HandleFunc(callbackUrl, callback)
}
