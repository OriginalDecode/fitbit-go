package fitbitAPI

import (
	"reflect"
)

type Scopes struct {
	Activity  string
	Heartrate string
	Location  string
	Nutrition string
	Profile   string
	Settings  string
	Sleep     string
	Social    string
	Weight    string
}

func GetScopes() Scopes {
	return Scopes{
		Activity:  "activity",
		Heartrate: "heartrate",
		Location:  "location",
		Nutrition: "nutrition",
		Profile:   "profile",
		Settings:  "settings",
		Sleep:     "sleep",
		Social:    "social",
		Weight:    "weight",
	}
}

func ValidScope(scope string) bool {
	scopes := GetScopes()
	s := reflect.ValueOf(&scopes).Elem()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		// fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
		if scope == f.Interface() {
			return true
		}

	}
	return false
}
