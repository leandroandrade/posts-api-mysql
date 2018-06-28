package service

import (
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"github.com/leandroandrade/posts-api-mysql/posts/model"
)

func (Service) Update(post model.Post) error {
	return processUpdate(post)
}

func processUpdate(post model.Post) error {
	query := fmt.Sprintf("UPDATE post SET description='%s' WHERE id=%d", post.Description, post.Id)
	_, err := mysql.DB.Exec(query)
	return err
}
