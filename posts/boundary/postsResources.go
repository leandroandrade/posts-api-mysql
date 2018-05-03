package boundary

import (
	"net/http"
	"github.com/leandroandrade/posts-api-mysql/posts/service"
	"encoding/json"
	"github.com/leandroandrade/posts-api-mysql/handler"
	"github.com/leandroandrade/posts-api-mysql/logger"
)

func GetPosts(writer http.ResponseWriter, _ *http.Request) *handler.AppError {
	posts, err := service.FindAll()
	if err != nil {
		logger.Error.Println(err.Error())
		return &handler.AppError{Error: err, Message: "Internal Error", Code: 500}
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(posts)
	writer.WriteHeader(http.StatusOK)

	return nil
}

func CreatePosts(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("CREATE POSTS"))
}
