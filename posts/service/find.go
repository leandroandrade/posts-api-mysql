package service

import (
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"errors"
	"github.com/leandroandrade/posts-api-mysql/posts/model"
)

func FindAll() ([]model.Post, error) {
	posts, err := getAllPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func getAllPosts() ([]model.Post, error) {
	rows, err := mysql.DB.Query("select * from post")
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	var posts []model.Post
	var post model.Post

	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.Description, &post.DatePosted); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, err
}
