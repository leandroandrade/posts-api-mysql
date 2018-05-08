package boundary

import (
	"net/http"
	"github.com/leandroandrade/posts-api-mysql/posts/service"
	"encoding/json"
	"github.com/leandroandrade/posts-api-mysql/handler"
	"github.com/leandroandrade/posts-api-mysql/logger"
	"io/ioutil"
	"strconv"
)

func GetPosts(writer http.ResponseWriter, _ *http.Request) *handler.AppError {
	posts, err := service.FindAll()
	if err != nil {
		logger.Error.Println(err.Error())
		return &handler.AppError{Error: err.Error(), Message: "internal Error", Code: 500}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(posts)

	return nil
}

func CreatePosts(writer http.ResponseWriter, request *http.Request) *handler.AppError {
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return &handler.AppError{Error: err.Error(), Message: "cannot read a content", Code: http.StatusBadRequest}
	}

	post, err := service.Save(body)
	if err != nil {
		return &handler.AppError{Error: err.Error(), Message: "cannot read a content", Code: http.StatusBadRequest}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Location", request.Host+request.URL.Path+"/"+strconv.Itoa(post.Id))

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(post)

	return nil
}
