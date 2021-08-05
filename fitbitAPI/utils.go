package fitbitAPI

import (
	"net/http"
	"os"
)

func enableCors(w *http.ResponseWriter) {
	env := os.Getenv("APP_ENV")
	if env == "DEBUG" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}
}
