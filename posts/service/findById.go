package service

import (
	"strconv"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"fmt"
)

func FindById(id string) (Post, error) {
	var post Post

	identifier, err := strconv.Atoi(id)
	if err != nil {
		return post, err
	}

	return post, getPostByID(identifier, &post)
}

func getPostByID(identifier int, post *Post) error {
	query := fmt.Sprintf("SELECT id, description, date_posted FROM post WHERE id=%d", identifier)
	return mysql.DB.QueryRow(query).Scan(&post.Id, &post.Description, &post.DatePosted)
}
