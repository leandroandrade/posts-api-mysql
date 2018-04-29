package service

import (
	"github.com/leandroandrade/posts-api-mysql/posts/model"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"log"
	"database/sql"
)

func FindAll() []model.Post {
	rows, err := getPosts()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var posts []model.Post
	var post model.Post

	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.Description); err != nil {
			log.Fatal(err.Error())
		}
		posts = append(posts, post)
	}
	return posts
}

func getPosts() (*sql.Rows, error) {
	rows, err := mysql.DB.Query("select * from post")
	return rows, err
}
