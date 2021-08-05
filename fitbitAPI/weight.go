package fitbitAPI

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Unmarshal(data []byte) ([]Sleep, error) {
	var tmp sleep
	err := json.Unmarshal(data, &tmp)
	return tmp.Sleep, err
}

const fatURL = "https://api.fitbit.com/1/user/-/body/log/fat/date"
const weightURL = "https://api.fitbit.com/1/user/-/body/log/weight/date/"

// https://api.fitbit.com/1/user/[user-id]/body/log/fat/date/[base-date]/[end-date].json

func RequestWeight(nowStr string, thenStr string) []byte {
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
