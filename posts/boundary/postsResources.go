package boundary

import (
	"net/http"
	"github.com/leandroandrade/posts-api-mysql/posts/service"
	"encoding/json"
	"github.com/leandroandrade/posts-api-mysql/handler"
	"github.com/leandroandrade/posts-api-mysql/logger"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/leandroandrade/posts-api-mysql/posts/model"
	"database/sql"
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
		logger.Error.Println(err.Error())
		return &handler.AppError{Error: err.Error(), Message: "cannot read a content", Code: http.StatusBadRequest}
	}

	post, err := service.Save(body)
	if err != nil {
		logger.Error.Println(err.Error())
		return &handler.AppError{Error: err.Error(), Message: "cannot read a content", Code: http.StatusBadRequest}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Location", request.Host+request.URL.Path+"/"+strconv.Itoa(post.Id))

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(post)

	return nil
}

func DeletePost(writer http.ResponseWriter, request *http.Request) *handler.AppError {
	vars := mux.Vars(request)

	err := service.DeleteByID(vars["id"])
	if err != nil {
		logger.Error.Println(err.Error())
		return &handler.AppError{Error: err.Error(), Message: "cannot remove the post", Code: http.StatusBadRequest}
	}

	writer.WriteHeader(http.StatusNoContent)

	return nil
}

func UpdatePost(writer http.ResponseWriter, request *http.Request) *handler.AppError {
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error.Println(err.Error())
		return &handler.AppError{Error: err.Error(), Message: "cannot update the post", Code: http.StatusBadRequest}
	}

	var post model.Post
	if err = json.NewDecoder(request.Body).Decode(&post); err != nil {
		logger.Error.Println(err.Error())
		return &handler.AppError{Error: err.Error(), Message: "cannot update the post", Code: http.StatusBadRequest}
	}

	post.Id = id
	if err = service.Update(&post); err != nil {
		logger.Error.Println(err.Error())
		return &handler.AppError{Error: err.Error(), Message: "cannot update the post", Code: http.StatusBadRequest}
	}

	writer.WriteHeader(http.StatusNoContent)

	return nil
}

func GetPostByID(writer http.ResponseWriter, request *http.Request) *handler.AppError {
	vars := mux.Vars(request)

	post, err := service.FindById(vars["id"])

	switch err {
	case sql.ErrNoRows:
		logger.Error.Println(err.Error())
		return &handler.AppError{Error: err.Error(), Message: "not found post id " + vars["id"], Code: http.StatusNotFound}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(post)

	return nil
}
