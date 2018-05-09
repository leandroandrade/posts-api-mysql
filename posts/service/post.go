package service

type Post struct {
	Id          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	DatePosted string `json:"date_posted,omitempty"`
}
