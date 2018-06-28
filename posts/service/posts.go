package service

import "github.com/leandroandrade/posts-api-mysql/posts/model"

type Posts interface {
	Save(body []byte) (*model.Post, error)
	DeleteByID(id string) error
	FindById(id string) (*model.Post, error)
	FindWithPagination(size string, page string) (*model.PostPaginationResponse, error)
	Update(post model.Post) error
}
