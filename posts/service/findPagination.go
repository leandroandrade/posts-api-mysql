package service

import (
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"errors"
	"fmt"
	"strconv"
	"math"
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
	total := getTotalDatabase()
	if total == 0 {
		return nil, errors.New("not exists values in database")
	}

	totalPages := int(math.Ceil(float64(total) / float64(size)))
	if page > totalPages {
		return nil, errors.New("number of page invalid")
	}

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

func getTotalDatabase() int {
	var total int
	query := fmt.Sprintf("SELECT COUNT(ID) FROM post")
	mysql.DB.QueryRow(query).Scan(&total)

	return total
}
