package auth

import (
	"github.com/leandroandrade/posts-api-mysql/jwt/keys"
	"golang.org/x/crypto/bcrypt"
	"github.com/leandroandrade/posts-api-mysql/authentication/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TODO alterar para buscar o usuario do DB
func Authenticate(user *model.User) bool {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	testUser := model.User{
		UUID:     user.UUID,
		Username: user.Username,
		Password: string(hashedPassword),
	}
	return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
}

func GenerateToken(userUUID string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(int(72))).Unix(),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}

	tokenString, err := token.SignedString(keys.PrivateKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}
