package service

import (
	"github.com/leandroandrade/posts-api-mysql/posts/model"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"database/sql"
	"errors"
)

func FindAll() ([]model.Post, error) {
	rows, err := getPosts()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	var posts []model.Post
	var post model.Post

	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.Description, &post.DateCreation); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPosts() (*sql.Rows, error) {
	rows, err := mysql.DB.Query("select * from post")
	return rows, err
}
