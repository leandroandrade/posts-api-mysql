package service

import (
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"time"
	"github.com/leandroandrade/posts-api-mysql/posts/model"
)

func (Service) Save(post *model.Post) error {
	if err := process(post); err != nil {
		return err
	}
	return nil

}

func process(post *model.Post) error {
	stmt, err := mysql.DB.Prepare("INSERT INTO post(description, date_posted) VALUES(?, ?)")
	if err != nil {
		return err
	}

	now := time.Now()
	post.DatePosted = now.Format("2006-01-02")
	result, err := stmt.Exec(post.Description, now.Local())
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
