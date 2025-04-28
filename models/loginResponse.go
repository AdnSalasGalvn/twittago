package models

type LoginResponse struct {
	Token string `json:"token,omitempty"` // token for the user
}
