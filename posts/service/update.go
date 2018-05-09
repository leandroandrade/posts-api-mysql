package service

import (
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/mysql"
)

func Update(post Post) error {
	return processUpdate(post)
}

func processUpdate(post Post) error {
	query := fmt.Sprintf("UPDATE post SET description='%s' WHERE id=%d", post.Description, post.Id)
	_, err := mysql.DB.Exec(query)
	return err
}
