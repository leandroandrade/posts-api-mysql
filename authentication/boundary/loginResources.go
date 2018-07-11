package boundary

import (
	"net/http"
	"github.com/leandroandrade/posts-api-mysql/authentication/model"
	"encoding/json"
	"github.com/leandroandrade/posts-api-mysql/logger"
	"fmt"
	"github.com/leandroandrade/posts-api-mysql/response"
	"github.com/leandroandrade/posts-api-mysql/authentication/service"
)

type LoginResources struct {
}

func NewLoginResources() *LoginResources {
	return &LoginResources{}
}

func (l LoginResources) Login(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		logger.Errorf("ERR: %v", err.Error())

		response.JSONErr(writer, response.Payload{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request",
			Detail:  fmt.Sprintf("cannot read a content: %v", err.Error()),
		})
		return
	}

	responseStatus, token := service.Login(&user)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(responseStatus)
	writer.Write(token)
}
