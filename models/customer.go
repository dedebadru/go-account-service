package models

import "time"

type Customer struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	IdentityNumber string    `json:"identity_number"`
	PhoneNumber    string    `json:"phone_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
