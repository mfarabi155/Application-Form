package models

import (
	"encoding/json"
	"time"
)

type Person struct {
	ID             uint      `json:"-"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	Name           string    `json:"name"`
	IdentityNumber string    `json:"identity_number"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
}

func (Person) TableName() string {
	return "persons"
}

func (p Person) MarshalJSON() ([]byte, error) {
	type Alias Person
	return json.Marshal(&struct {
		Alias
		DateOfBirth string `json:"date_of_birth"`
	}{
		Alias:       (Alias)(p),
		DateOfBirth: p.DateOfBirth.Format("2006-01-02"),
	})
}
