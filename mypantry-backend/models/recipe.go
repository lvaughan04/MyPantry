package models

import "time"

type Recipe struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Author      string       `json:"author"`
	PrepTime    int          `json:"prep_time"` //in minutes
	CookTime    int          `json:"cook_time"` //in minutes
	Servings    int          `json:"servings"`
	Description string       `json:"description"`
	Instruction []string     `json:"instruction"`
	Ingredients []Ingredient `json:"ingredients"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
