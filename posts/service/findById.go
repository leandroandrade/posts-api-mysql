package service

import (
	"strconv"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/posts/model"
	"database/sql"
	"errors"
)

func (Service) FindById(id string) (*model.Post, error) {
	var post model.Post

	identifier, err := strconv.Atoi(id)
	if err != nil {
		return &post, err
	}

	err = getPostByID(identifier, &post)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, errors.New("post not found")
		default:
			return nil, fmt.Errorf("cannot find post: %v", err.Error())
		}
	}

	return &post, err
}

func getPostByID(identifier int, post *model.Post) error {
	query := fmt.Sprintf("SELECT id, description, date_posted FROM post WHERE id=%d", identifier)
	return mysql.DB.QueryRow(query).Scan(&post.Id, &post.Description, &post.DatePosted)
}
