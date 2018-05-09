package service

import (
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"strconv"
	"errors"
)

func DeleteByID(id string) error {
	identifier, _ := strconv.Atoi(id)
	if err := verifyPostExists(identifier); err != nil {
		return err
	}

	return deleteFromDatabase(identifier)
}

func verifyPostExists(identifier int) error {
	var exists bool

	query := fmt.Sprintf("SELECT IF(COUNT(id),'true','false') FROM post WHERE id=%d", identifier)
	mysql.DB.QueryRow(query).Scan(&exists)
	if !exists {
		return errors.New(fmt.Sprintf("the post %v not exist!", identifier))
	}
	return nil

}

func deleteFromDatabase(identifier int) error {
	query := fmt.Sprintf("delete from post where id=%d", identifier)
	_, err := mysql.DB.Exec(query)
	return err
}
