package fitbitData

import (
	"encoding/json"
)

const ProfileUrl = "/1/user/-/profile.json"

type Badge struct {
	GradientEndColor   string      `json:"badgeGradientEndColor"`
	GradientStartColor string      `json:"badgegradientStartColor"`
	BadgeType          string      `json:"badgeType"`
	Category           string      `json:"category"`
	Cheers             interface{} `json:"cheers"`
	Date               string      `json:"dateTime"`
	Desc               string      `json:"description"`
	Msg                string      `json:"earnedMessage"`
	EncodedId          string      `json:"encodedId"`
	Image300px         string      `json:"300px"`
	Image125px         string      `json:"125px"`
	Image100px         string      `json:"100px"`
	Image75px          string      `json:"75px"`
	Image50px          string      `json:"50px"`
	MarketingDesc      string      `json:"marketingDescription"`
	MobileDesc         string      `json:"mobileDescription"`
	Name               string      `json:"name"`
	ShareImage640px    string      `json:"shareImage640px"`
	ShareText          string      `json:"shareText"`
	ShortDescription   string      `json:"shortDescription"`
	ShortName          string      `json:"shortName"`
	TimesAchieved      int32       `json:"timesAchieved"`
	Value              int32       `json:"value"`
}

type UserData struct {
	Age                     int32       `json:"age"`
	IsAmbassador            bool        `json:"ambassador"`
	AvatarUrl               string      `json:"avatar"`
	Avatar150               string      `json:"avatar150"`
	Avatar640               string      `json:"avatar640"`
	AvgDailySteps           int32       `json:"averageDailySteps"`
	ClockDisplayFormat      string      `json:"clockTimeDisplayFormat"`
	Corporate               bool        `json:"corporate"`
	CorporateAdmin          bool        `json:"corporateAdmin"`
	Country                 string      `json:"country"`
	DateOfBirth             string      `json:"dateOfBirth"`
	DisplayName             string      `json:"displayName"`
	DisplayNameSetting      string      `json:"name"`
	DistanceUnit            string      `json:"distanceUnit"`
	EncodedId               string      `json:"encodedId"`
	FamilyGuidanceEnabled   bool        `json:"familyGuidanceEnabled"`
	Features                interface{} `json:"features"`
	FirstName               string      `json:"firstName"`
	FoodsLocale             string      `json:"foodsLocale"`
	FullName                string      `json:"fullName"`
	Gender                  string      `json:"gender"`
	GlucoseUnit             string      `json:"glucoseUnit"`
	Height                  int32       `json:"height"`
	HeightUnit              string      `json:"heightUnit"`
	IsChild                 bool        `json:"isChild"`
	IsCoach                 bool        `json:"isCoach"`
	LastName                string      `json:"lastName"`
	Locale                  string      `json:"locale"`
	MemberSince             string      `json:"memberSince"`
	MfaEnabled              bool        `json:"mfaEnabled"`
	OffsetFromUTCMillis     int32       `json:"offsetFromUTCMillis"`
	StartDayOfWeek          string      `json:"startDayOfWeek"`
	StrideLengthRunning     float32     `json:"strideLengthRunning"`
	StrideLengthRunningType string      `json:"strideLengthRunningType"`
	StrideLengthWalking     float32     `json:"strideLengthWalking"`
	StrideLengthWalkingType string      `json:"strideLengthWalkingType"`
	SwimUnit                string      `json:"swimUnit"`
	Timezone                string      `json:"timezone"`
	TopBadges               []Badge     `json:"topBadges"`
	WaterUnit               string      `json:"waterUnit"`
	WaterUnitName           string      `json:"waterUnitName"`
	Weight                  int32       `json:"weight"`
	WeightUnit              string      `json:"weightUnit"`
}

type User struct {
	Profile UserData `json:"user"`
}

func UnmarshalProfile(data []byte) (UserData, error) {
	var user User
	err := json.Unmarshal(data, &user)

	return user.Profile, err
}
