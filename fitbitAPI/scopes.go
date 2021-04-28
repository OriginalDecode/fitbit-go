package fitbitAPI

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
