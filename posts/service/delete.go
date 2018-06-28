package service

import (
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"strconv"
)

func (Service) DeleteByID(id string) error {
	identifier, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("the 'id' is not a number: %v", err.Error())
	}

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
		return fmt.Errorf("the post %v not exist", identifier)
	}
	return nil

}

func deleteFromDatabase(identifier int) error {
	query := fmt.Sprintf("delete from post where id=%d", identifier)
	_, err := mysql.DB.Exec(query)
	return err
}
