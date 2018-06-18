package model

import "time"

type Post struct {
	Id          int       `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	DatePosted  time.Time `json:"date_posted,omitempty"`
}
