package boundary

import (
	"net/http"
	"github.com/leandroandrade/posts-api-mysql/posts/service"
	"encoding/json"
	"github.com/leandroandrade/posts-api-mysql/logger"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
	"database/sql"
	"github.com/leandroandrade/posts-api-mysql/posts/model"
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/response"
)

func GetPosts(writer http.ResponseWriter, _ *http.Request) {
	posts, err := service.FindAll()
	if err != nil {
		logger.Error(err.Error())

		response.JSON(writer, response.Message{
			Code:             http.StatusInternalServerError,
			MessageUser:      "Cannot get post",
			MessageDeveloper: fmt.Sprintf("internal Error: %v", err.Error()),
		})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(posts)
}

func CreatePosts(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		logger.Error(err.Error())

		response.JSON(writer, response.Message{
			Code:             http.StatusBadRequest,
			MessageUser:      "Cannot save post",
			MessageDeveloper: fmt.Sprintf("cannot read a content: %v", err.Error()),
		})
		return
	}

	post, err := service.Save(body)
	if err != nil {
		logger.Error(err.Error())

		response.JSON(writer, response.Message{
			Code:             http.StatusBadRequest,
			MessageUser:      "Cannot save post",
			MessageDeveloper: fmt.Sprintf("cannot read a content: %v", err.Error()),
		})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Location", request.Host+request.URL.Path+"/"+strconv.Itoa(post.Id))

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(post)
}

func DeletePost(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	err := service.DeleteByID(vars["id"])
	if err != nil {
		logger.Error(err.Error())

		response.JSON(writer, response.Message{
			Code:             http.StatusBadRequest,
			MessageUser:      "Cannot delete post",
			MessageDeveloper: fmt.Sprintf("cannot remove the post: %v", err.Error()),
		})
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func UpdatePost(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id, _ := strconv.Atoi(vars["id"])

	var post model.Post
	if err := json.NewDecoder(request.Body).Decode(&post); err != nil {
		logger.Error(err.Error())

		response.JSON(writer, response.Message{
			Code:             http.StatusBadRequest,
			MessageUser:      "Cannot update post",
			MessageDeveloper: fmt.Sprintf("cannot update the post: %v", err.Error()),
		})
		return
	}

	post.Id = id
	if err := service.Update(post); err != nil {
		logger.Error(err.Error())

		response.JSON(writer, response.Message{
			Code:             http.StatusBadRequest,
			MessageUser:      "Cannot update post",
			MessageDeveloper: fmt.Sprintf("cannot update the post: %v", err.Error()),
		})
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func GetPostByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	post, err := service.FindById(vars["id"])

	switch err {
	case sql.ErrNoRows:
		logger.Error(err.Error())

		response.JSON(writer, response.Message{
			Code:             http.StatusNotFound,
			MessageUser:      "Failed to get post",
			MessageDeveloper: err.Error(),
		})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(post)
}

func FindPostsPagination(writer http.ResponseWriter, request *http.Request) {
	size := request.URL.Query().Get("size")
	page := request.URL.Query().Get("page")

	posts, err := service.FindWithPagination(size, page)
	if err != nil {
		logger.Error(err.Error())

		response.JSON(writer, response.Message{
			Code:             http.StatusBadRequest,
			MessageUser:      "Failed when list posts",
			MessageDeveloper: err.Error(),
		})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(posts)
}
