package fitbitData

import (
	"encoding/json"
)

type ResponseError struct {
	ErrorType string `json:"errorType"`
	Message   string `json:"message"`
}

type AuthResponse struct {
	AccessToken  string          `json:"access_token"`
	ExpiresIn    int32           `json:"expires_in"`
	RefreshToken string          `json:"refresh_token"`
	Scopes       string          `json:"scope"`
	TokenType    string          `json:"token_type"`
	UserId       string          `json:"user_id"`
	Success      bool            `json:"success"`
	Errors       []ResponseError `json:"errors"`
}

func UnmarshalAuth(data []byte) (AuthResponse, error) {
	var auth AuthResponse
	err := json.Unmarshal(data, &auth)
	return auth, err
}
