package service

import (
	"github.com/leandroandrade/posts-api-mysql/posts/model"
	"strconv"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"fmt"
)

func FindById(id string) (model.Post, error) {
	var post model.Post

	identifier, err := strconv.Atoi(id)
	if err != nil {
		return post, err
	}

	return post, getPostByID(identifier, &post)
}

func getPostByID(identifier int, post *model.Post) error {
	query := fmt.Sprintf("SELECT id, description, date_creation FROM post WHERE id=%d", identifier)
	return mysql.DB.QueryRow(query).Scan(&post.Id, &post.Description, &post.DateCreation)
}
