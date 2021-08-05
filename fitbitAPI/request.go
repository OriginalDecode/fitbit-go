package fitbitAPI

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func parseTime(time time.Time) string {
	year := time.Year()
	month := int(time.Month())
	day := time.Day()

	if month < 10 {
		if day < 10 {
			return fmt.Sprintf("%d-0%d-0%d", year, month, day)
		}
		return fmt.Sprintf("%d-0%d-%d", year, month, day)
	}
	if day < 10 {
		return fmt.Sprintf("%d-%d-0%d", year, month, day)
	}
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

func monthsCountSince(createdAtTime time.Time) int {
	now := time.Now()
	months := 0
	month := createdAtTime.Month()
	for createdAtTime.Before(now) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextMonth := createdAtTime.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	return months
}

func DoRequest(from time.Time, to time.Time) {
	// RequestWeight(parseTime(to), parseTime(from))
}

func Request(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	query := r.URL.Query()

	for i, _ := range query {
		if ValidScope(i) == true {
			fmt.Printf("%s - ValidScope\n", i)

			// value, _ := time.Parse("2018-01-01", "2018-01-01")

			location, _ := time.LoadLocation("Europe/Rome")
			start := time.Date(2018, 01, 01, 0, 0, 0, 0, location)

			months := monthsCountSince(start)

			for i := 0; i < months; i++ {
				DoRequest(start.AddDate(0, i, 0), start.AddDate(0, i+1, 0))
			}

			// now := time.Now()
			// then := now.AddDate(0, -1, 0) // get a month back
			// RequestWeight(parseTime(now), parseTime(then))

		} else {
			log.Println("Error, invalid scope")
			return
		}
	}

}
