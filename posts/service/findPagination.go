package service

import (
	"github.com/leandroandrade/posts-api-mysql/mysql"
	"errors"
	"fmt"
	"strconv"
	"math"
	"github.com/leandroandrade/posts-api-mysql/posts/model"
)

const maxsize = 50

func (Service) FindWithPagination(size string, page string) (*model.PostPaginationResponse, error) {
	sizeResult, err := strconv.Atoi(size)
	if err != nil {
		return nil, fmt.Errorf("field 'size' invalid: it is not a number: %v", err.Error())
	}

	pageResult, err := strconv.Atoi(page)
	if err != nil {
		return nil, fmt.Errorf("field 'page' invalid: it is not a number: %v", err.Error())
	}

	if pageResult <= 0 {
		return nil, fmt.Errorf("field 'page' needs to be greater than or equal 1")
	}

	if sizeResult > maxsize {
		return nil, fmt.Errorf("field 'size' allow max 50 elements for each page")
	}

	posts, err := getPostsPagination(sizeResult, pageResult)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func getPostsPagination(size int, page int) (*model.PostPaginationResponse, error) {
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

	posts := make([]*model.Post, 0)

	for rows.Next() {
		var post model.Post
		if err = rows.Scan(&post.Id, &post.Description, &post.DatePosted); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return model.New(posts, total, totalPages, len(posts), page), nil
}

func getTotalDatabase() int {
	var total int
	query := fmt.Sprintf("SELECT COUNT(ID) FROM post")
	mysql.DB.QueryRow(query).Scan(&total)

	return total
}
