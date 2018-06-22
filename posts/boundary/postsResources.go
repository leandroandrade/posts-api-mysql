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

func CreatePosts(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		logger.Error(err.Error())

		response.JSONErr(writer, response.Payload{
			Code:    http.StatusBadRequest,
			Message: "Cannot save post",
			Detail:  fmt.Sprintf("cannot read a content: %v", err.Error()),
		})
		return
	}

	post, err := service.Save(body)
	if err != nil {
		logger.Error(err.Error())

		response.JSONErr(writer, response.Payload{
			Code:    http.StatusBadRequest,
			Message: "Cannot save post",
			Detail:  fmt.Sprintf("cannot read a content: %v", err.Error()),
		})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Location", request.Host+request.URL.Path+"/"+strconv.Itoa(post.Id))

	response.JSON(writer, http.StatusCreated, post)
}

func DeletePost(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	err := service.DeleteByID(vars["id"])
	if err != nil {
		logger.Error(err.Error())

		response.JSONErr(writer, response.Payload{
			Code:    http.StatusBadRequest,
			Message: "Cannot delete post",
			Detail:  fmt.Sprintf("cannot remove the post: %v", err.Error()),
		})
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func UpdatePost(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.JSONErr(writer, response.Payload{
			Code:    http.StatusBadRequest,
			Message: "Cannot update post",
			Detail:  fmt.Sprintf("cannot update the post: field 'id' is not a number: %v", err.Error()),
		})
		return
	}

	var post model.Post
	if err := json.NewDecoder(request.Body).Decode(&post); err != nil {
		logger.Error(err.Error())

		response.JSONErr(writer, response.Payload{
			Code:    http.StatusBadRequest,
			Message: "Cannot update post",
			Detail:  fmt.Sprintf("cannot update the post: %v", err.Error()),
		})
		return
	}

	post.Id = id
	if err := service.Update(post); err != nil {
		logger.Error(err.Error())

		response.JSONErr(writer, response.Payload{
			Code:    http.StatusBadRequest,
			Message: "Cannot update post",
			Detail:  fmt.Sprintf("cannot update the post: %v", err.Error()),
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

		response.JSONErr(writer, response.Payload{
			Code:    http.StatusNotFound,
			Message: "Failed to get post",
			Detail:  err.Error(),
		})
		return
	}

	response.JSON(writer, http.StatusOK, post)
}

func FindPostsPagination(writer http.ResponseWriter, request *http.Request) {
	size := request.URL.Query().Get("size")
	page := request.URL.Query().Get("page")

	posts, err := service.FindWithPagination(size, page)
	if err != nil {
		logger.Error(err.Error())

		response.JSONErr(writer, response.Payload{
			Code:    http.StatusBadRequest,
			Message: "Failed when list posts",
			Detail:  err.Error(),
		})
		return
	}

	response.JSON(writer, http.StatusOK, posts)
}
