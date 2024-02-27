package domain

type User struct {
	UserID string `json:"use_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
