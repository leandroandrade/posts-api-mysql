package service

import (
	"encoding/json"
	"errors"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"time"
)

func Save(body []byte) (Post, error) {
	var post Post
	if err := json.Unmarshal(body, &post); err != nil {
		return post, errors.New("cannot unmarshal content")
	}

	if err := process(&post); err != nil {
		return post, err
	}

	return post, nil

}

func process(post *Post) error {
	stmt, err := mysql.DB.Prepare("INSERT INTO post(description, date_posted) VALUES(?, ?)")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(post.Description, time.Now().Local())
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	post.Id = int(lastID)
	return nil

}
