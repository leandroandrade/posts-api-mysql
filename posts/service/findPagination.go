package service

import (
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"errors"
	"fmt"
	"strconv"
)

func FindWithPagination(size string, page string) ([]Post, error) {
	sizeResult, _ := strconv.Atoi(size)
	pageResult, _ := strconv.Atoi(page)

	posts, err := getPostsPagination(sizeResult, pageResult)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func getPostsPagination(size int, page int) ([]Post, error) {
	pageNumber := size * (page - 1)

	query := fmt.Sprintf("SELECT * FROM post LIMIT %d OFFSET %d", size, pageNumber)
	rows, err := mysql.DB.Query(query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	var posts []Post
	var post Post

	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.Description, &post.DatePosted); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, err
}
