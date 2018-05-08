package service

import (
	"github.com/leandroandrade/posts-api-mysql/posts/model"
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/mysql"
)

func Update(post *model.Post) error {
	return processUpdate(*post)
}

func processUpdate(post model.Post) error {
	statement := fmt.Sprintf("UPDATE post SET description='%s' WHERE id=%d", post.Description, post.Id)
	_, err := mysql.DB.Exec(statement)
	return err
}
