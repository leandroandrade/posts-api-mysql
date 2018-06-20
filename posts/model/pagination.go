package model

type PostPaginationResponse struct {
	Content       []*Post `json:"content,omitempty"`
	TotalElements int     `json:"totalElements,omitempty"`
	TotalPages    int     `json:"totalPages,omitempty"`
	Size          int     `json:"size,omitempty"`
	Page          int     `json:"page,omitempty"`
}

func New(posts []*Post, totalElements int, totalPages int, size int, page int) *PostPaginationResponse {
	return &PostPaginationResponse{
		Content:       posts,
		TotalElements: totalElements,
		TotalPages:    totalPages,
		Size:          size,
		Page:          page,
	}
}
