package service

import (
	"github.com/leandroandrade/posts-api-mysql/authentication/model"
	"github.com/leandroandrade/posts-api-mysql/jwt/auth"
	"net/http"
	"encoding/json"
	"github.com/leandroandrade/posts-api-mysql/jwt/token"
)

func Login(user *model.User) (int, []byte) {
	authentication := auth.InitJWTAuthenticationProcess()

	if !authentication.Authenticate(user) {
		return http.StatusUnauthorized, []byte("")
	}

	tokenauth, err := authentication.GenerateToken(user.UUID)
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	}

	response, _ := json.Marshal(token.JWTToken{Token: tokenauth})
	return http.StatusOK, response
}
