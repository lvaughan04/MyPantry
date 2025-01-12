package models

import "time"

// type of ingredients: spice, meat, diary, grains, condements
type Ingredient struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Quantity       string    `json:"quantity"`
	Category       string    `json:"category"`
	ExpirationDate time.Time `json:"expiration_date"`
}
