package service

import (
	"github.com/leandroandrade/posts-api-mysql/posts/model"
)

type ServiceMock struct{}

func NewServiceMock() *ServiceMock {
	return &ServiceMock{}
}

func (ServiceMock) Save(post *model.Post) error {
	post.Id = 9988
	return nil
}

func (ServiceMock) DeleteByID(id string) error {
	return nil
}

func (ServiceMock) FindById(id string) (*model.Post, error) {
	return nil, nil
}

func (ServiceMock) FindWithPagination(size string, page string) (*model.PostPaginationResponse, error) {
	return nil, nil
}

func (ServiceMock) Update(post model.Post) error {
	return nil
}
