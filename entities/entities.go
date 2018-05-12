package entities

import (
	"time"
)

// Passport holds passport data
type Passport struct {
	ID           string    `json:"id"`
	DateOfIssue  time.Time `json:"dateOfIssue"`
	DateOfExpiry time.Time `json:"dateOfExpiry"`
	Authority    string    `json:"authority"`
	UserID       int       `json:"userId"`
}

// User holds personal user information
// swagger:response user
type User struct {
	// UID
	ID int `json:"id"`
	// First name
	FirstName string `json:"firstName"`
	// Last name
	LastName string `json:"lastName"`
	// Date of birth
	DateOfBirth time.Time `json:"dateOfBirth"`
	// Location of birth
	LocationOfBirth string `json:"locationOfBirth"`
}
