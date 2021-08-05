package fitbitAPI

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const sleepURL = "https://api.fitbit.com/1.2/user/-/sleep/date/"

type SleepStage struct {
	Count            int32 `json:"count"`
	Minutes          int32 `json:"minutes"`
	ThirtyDayAvgMins int32 `json:"thirtyDayAvgMinutes"`
}

type SleepSummary struct {
	Deep  SleepStage `json:"deep"`
	Light SleepStage `json:"light"`
	Rem   SleepStage `json:"rem"`
	Wake  SleepStage `json:"wake"`
}

type SleepPoint struct {
	DateTime string `json:"dateTime"`
	Level    string `json:"level"`
	Seconds  int32  `json:"seconds"`
}

type SleepData struct {
	LongData  []SleepPoint `json:"data"`
	ShortData []SleepPoint `json:"shortData"`
	Summary   SleepSummary `json:"summary"`
}

type Sleep struct {
	Date                string    `json:"dateOfSleep"`
	Duration            int32     `json:"duration"` //in seconds
	Efficiency          int32     `json:"efficiency"`
	EndTime             string    `json:"endTime"`
	InfoCode            int32     `json:"infoCode"`
	IsMainSleep         bool      `json:"isMainSleep"`
	Levels              SleepData `json:"levels"`
	LogId               int64     `json:"logId"`
	MinutesAfterWakeup  int32     `json:"minutesAfterWakeup"`
	MinutesAsleep       int32     `json:"minutesAsleep"`
	MinutesAwake        int32     `json:"minutesAwake"`
	MinutesToFallAsleep int32     `json:"minutesToFallAsleep"`
	StartTime           string    `json:"startTime"`
	TimeInBed           int32     `json:"timeInBed"`
	Type                string    `json:"type"`
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
type sleep struct {
	Sleep []Sleep `json:"sleep"`
}

func UnmarshalSleep(data []byte) ([]Sleep, error) {
	var tmp sleep
	err := json.Unmarshal(data, &tmp)
	return tmp.Sleep, err
}

func RequestSleep(nowStr string, thenStr string) []byte {
	url := weightURL + thenStr + "/" + nowStr + ".json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
	}

	client := &http.Client{Timeout: time.Second * 10}

	req.Header.Set("Authorization", "Bearer "+GetAccessToken())

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	return body
}

// func handleSleepRequest(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Println(r.RequestURI)
// 	// fmt.Printf("[%d] %s\n", time.Now().Unix(), "Sleep Request")

// 	// defer r.Body.Close()
// 	// body, err := ioutil.ReadAll(r.Body)
// 	// if err != nil {
// 	// 	fmt.Println("{}", err.Error())
// 	// }

// 	// var data SleepRequest
// 	// err = schema.NewDecoder().Decode(&data, r.URL.Query())
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }

// 	// url := fitbitData.SleepUrl + data.StartDate + "/" + data.EndDate + ".json"
// 	// fmt.Println(url)
// 	// req, err := http.NewRequest("GET", url, nil)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }

// 	// req.Header.Set("Authorization", "Bearer "+RespData.AccessToken)
// 	// resp, err := CLIENT.Do(req)
// 	// if err != nil {
// 	// 	fmt.Println("{}", err.Error())
// 	// }

// 	// defer resp.Body.Close()

// 	// body, err = ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	fmt.Println("{}", err.Error())
// 	// }

// 	// sleep, err := fitbitData.UnmarshalSleep(body)
// 	// if err != nil {
// 	// 	fmt.Println("{}", err.Error())
// 	// }

// 	// // Creating the sleep data response
// 	// w.Header().Set("Content-Type", "application/json")
// 	// w.Header().Set("Access-Control-Allow-Origin", "*")

// 	// /* Setup the last 7 days */

// 	// /*
// 	// 	type SleepStage struct {
// 	// 		Count            int32 `json:"count"`
// 	// 		Minutes          int32 `json:"minutes"`
// 	// 		ThirtyDayAvgMins int32 `json:"thirtyDayAvgMinutes"`
// 	// 	}
// 	// */
// 	// sleepResponse := SleepResponse{}
// 	// for i := len(sleep) - 1; i >= 0; i-- {
// 	// 	summary := &sleep[i].Levels.Summary
// 	// 	sleepResponse.Dates = append(sleepResponse.Dates, sleep[i].Date)
// 	// 	sleepResponse.DeepSleep = append(sleepResponse.DeepSleep, summary.Deep.Minutes)
// 	// 	sleepResponse.LightSleep = append(sleepResponse.LightSleep, summary.Light.Minutes)
// 	// 	sleepResponse.RemSleep = append(sleepResponse.RemSleep, summary.Rem.Minutes)
// 	// 	sleepResponse.Wake = append(sleepResponse.Wake, summary.Wake.Minutes)
// 	// }

// 	// json.NewEncoder(w).Encode(sleepResponse)

// 	// var sleepDuration int32 = 0
// 	// var longData = &sleep[0].Levels.LongData
// 	// for i := 0; i < len(*longData); i++ {
// 	// 	level := (*longData)[i].Level
// 	// 	if level == "light" || level == "deep" || level == "rem" {
// 	// 		sleepDuration += (*longData)[i].Seconds
// 	// 	}
// 	// }
// 	// fmt.Println(sleepDuration)
// 	// fmt.Println(sleep[0].Date)
// 	// fmt.Println(sleep[0].Duration)
// 	// fmt.Println(sleep[0].Efficiency)
// 	// fmt.Println(sleep[0].EndTime)
// 	// fmt.Println(sleep[0].InfoCode)

// }
