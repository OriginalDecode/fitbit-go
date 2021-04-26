package fitbitData

import (
	"encoding/json"
)

const SleepUrl = "https://api.fitbit.com/1.2/user/-/sleep/date/"

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

type _Sleep struct {
	Sleep []Sleep `json:"sleep"`
}

func UnmarshalSleep(data []byte) ([]Sleep, error) {
	var tmp _Sleep
	err := json.Unmarshal(data, &tmp)

	return tmp.Sleep, err
}
