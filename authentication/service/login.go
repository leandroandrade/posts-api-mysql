package service

import (
	"github.com/leandroandrade/posts-api-mysql/authentication/model"
	"github.com/leandroandrade/posts-api-mysql/jwt/auth"
	"net/http"
	"encoding/json"
)

func Login(user *model.User) (int, []byte) {
	if !auth.Authenticate(user) {
		return http.StatusUnauthorized, []byte("")
	}

	tokenauth, err := auth.GenerateToken(user.UUID)
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	}

	response, _ := json.Marshal(model.Token{Token: tokenauth})
	return http.StatusOK, response
}
