package auth

import (
	"crypto/rsa"
	"github.com/leandroandrade/posts-api-mysql/jwt/keys"
	"golang.org/x/crypto/bcrypt"
	"github.com/leandroandrade/posts-api-mysql/authentication/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTAuthenticationProcess struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var jwtInstance *JWTAuthenticationProcess = nil

func InitJWTAuthenticationProcess() *JWTAuthenticationProcess {
	if jwtInstance == nil {
		jwtInstance = &JWTAuthenticationProcess{
			privateKey: keys.GetPrivateKey(),
			PublicKey:  keys.GetPublicKey(),
		}
	}
	return jwtInstance
}

// TODO alterar para buscar o usuario do DB
func (j *JWTAuthenticationProcess) Authenticate(user *model.User) bool {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	testUser := model.User{
		UUID:     user.UUID,
		Username: user.Username,
		Password: string(hashedPassword),
	}

	return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
}

func (j *JWTAuthenticationProcess) GenerateToken(userUUID string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(int(72))).Unix(),
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}
	tokenString, err := token.SignedString(j.privateKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}
