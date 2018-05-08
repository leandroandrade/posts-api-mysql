package service

import (
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"strconv"
)

func DeleteByID(id string) error {
	identifier, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return deleteFromDatabase(identifier)
}

func deleteFromDatabase(identifier int) error {
	query := fmt.Sprintf("delete from post where id=%d", identifier)
	_, err := mysql.DB.Exec(query)
	return err
}
