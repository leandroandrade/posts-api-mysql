package model

type Post struct {
	Id          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	DateCreation string `json:"date_creation,omitempty"`
}
