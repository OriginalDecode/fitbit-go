package main

import (
	"./fitbitAPI"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/fitbit"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
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

var (
	CLIENT_ID     string = "WILL_BE_REPLACED"
	CLIENT_SECRET string = "WILL_BE_REPLACED"
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

var RespData fitbitData.AuthResponse

var CLIENT = http.DefaultClient

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
<body>
	<a href="/get/login">Fitbit Log In</a>
</body>
</html>`
	fmt.Fprintf(w, htmlIndex)
}

var oauth2Config *oauth2.Config

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("auth")
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func requestData(w http.ResponseWriter, r *http.Request) {

	// req, err := http.NewRequest("GET", API_URL+fitbitData.RequestUrl, nil)
	// if err != nil {
	// 	fmt.Println("{}", err)
	// }

	// req.Header.Set("Authorization", "Bearer "+AuthData.AccessToken)
	// resp, err := CLIENT.Do(req)
	// if err != nil {
	// 	fmt.Println("{}", err)
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("{}", err)
	// }

	// profile, err := fitbitData.UnmarshalProfile(body)
	// if err != nil {
	// 	fmt.Println("{}", err)
	// }

	// s := fmt.Sprintf("%d", profile.Age)

	// var htmlIndex = `<html>
	// <body>
	// 	age : ` + s + `</body>
	// </html>`
	// fmt.Fprintf(w, htmlIndex)

}

func handleCallback(w http.ResponseWriter, r *http.Request) {

	// requestBody := url.Values{}
	// requestBody.Set("code", AUTH_CODE)
	// requestBody.Set("grant_type", "authorization_code")
	// requestBody.Set("client_id", CLIENT_ID)
	// requestBody.Set("redirect_uri", "http://127.0.0.1:8080/callback")

	// req, err := http.NewRequest("POST", TOKEN_URL, strings.NewReader(requestBody.Encode()))
	// if err != nil {
	// 	panic(err.Error())
	// }

	// key := base64.StdEncoding.EncodeToString([]byte(CLIENT_ID + ":" + CLIENT_SECRET))

	// req.Header.Set("Authorization", "Basic "+key)
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// resp, err := CLIENT.Do(req)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer resp.Body.Close()
	// fmt.Println("Request Sent...")

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// AuthData, err = fitbitData.UnmarshalAuth([]byte(body))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// http.Redirect(w, r, "http://127.0.0.1:8081/access", 200)

	// var htmlIndex = `<html>
	// <body>
	// 	<p>logged in</p>
	// 	<a href="/request">Request Data</a>
	// </body>
	// </html>`
	// fmt.Fprintf(w, htmlIndex)
}

type WebRequest struct {
	RequestType string `json:"request"`
}

type SleepRequest struct {
	StartDate string `schema:"startDate"`
	EndDate   string `schema:"endDate"`
}

type SleepResponse struct {
	Dates      []string `json:"dates"`
	DeepSleep  []int32  `json:"deepSleep"`
	LightSleep []int32  `json:"lightSleep"`
	RemSleep   []int32  `json:"remSleep"`
	Wake       []int32  `json:"wake"`
}

func handleSleepRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	fmt.Printf("[%d] %s\n", time.Now().Unix(), "Sleep Request")

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("{}", err.Error())
	}

	var data SleepRequest
	err = schema.NewDecoder().Decode(&data, r.URL.Query())
	if err != nil {
		fmt.Println(err.Error())
	}

	url := fitbitData.SleepUrl + data.StartDate + "/" + data.EndDate + ".json"
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Set("Authorization", "Bearer "+RespData.AccessToken)
	resp, err := CLIENT.Do(req)
	if err != nil {
		fmt.Println("{}", err.Error())
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("{}", err.Error())
	}

	sleep, err := fitbitData.UnmarshalSleep(body)
	if err != nil {
		fmt.Println("{}", err.Error())
	}

	// Creating the sleep data response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	/* Setup the last 7 days */

	/*
		type SleepStage struct {
			Count            int32 `json:"count"`
			Minutes          int32 `json:"minutes"`
			ThirtyDayAvgMins int32 `json:"thirtyDayAvgMinutes"`
		}
	*/
	sleepResponse := SleepResponse{}
	for i := len(sleep) - 1; i >= 0; i-- {
		summary := &sleep[i].Levels.Summary
		sleepResponse.Dates = append(sleepResponse.Dates, sleep[i].Date)
		sleepResponse.DeepSleep = append(sleepResponse.DeepSleep, summary.Deep.Minutes)
		sleepResponse.LightSleep = append(sleepResponse.LightSleep, summary.Light.Minutes)
		sleepResponse.RemSleep = append(sleepResponse.RemSleep, summary.Rem.Minutes)
		sleepResponse.Wake = append(sleepResponse.Wake, summary.Wake.Minutes)
	}

	json.NewEncoder(w).Encode(sleepResponse)

	// var sleepDuration int32 = 0
	// var longData = &sleep[0].Levels.LongData
	// for i := 0; i < len(*longData); i++ {
	// 	level := (*longData)[i].Level
	// 	if level == "light" || level == "deep" || level == "rem" {
	// 		sleepDuration += (*longData)[i].Seconds
	// 	}
	// }
	// fmt.Println(sleepDuration)
	// fmt.Println(sleep[0].Date)
	// fmt.Println(sleep[0].Duration)
	// fmt.Println(sleep[0].Efficiency)
	// fmt.Println(sleep[0].EndTime)
	// fmt.Println(sleep[0].InfoCode)

}

func handleAuth(w http.ResponseWriter, r *http.Request) {

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

	fmt.Println(data.Code)

	requestBody := url.Values{}
	requestBody.Set("code", data.Code) /* ??? */
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
	fmt.Println("Sent Authorization")

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	RespData, err = fitbitData.UnmarshalAuth([]byte(body))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	CLIENT_ID = os.Getenv("FITBITAPI_CLIENT_ID")
	CLIENT_SECRET = os.Getenv("FITBITAPI_CLIENT_SECRET")

	oauth2Config = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:8080/callback",
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		Endpoint:     fitbit.Endpoint,
		Scopes: []string{ScopeActivity, ScopeHeartrate, ScopeLocation, ScopeNutrition, ScopeProfile,
			ScopeSettings, ScopeSleep, ScopeSocial, ScopeWeight},
	}

	fmt.Println("Running...")

	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/", fs)

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/post/auth", handleAuth)
	http.HandleFunc("/get/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/request", requestData)
	http.HandleFunc("/get/sleep", handleSleepRequest)
	http.ListenAndServe(":8080", nil)

}
