package token

type JWTToken struct {
	Token string `json:"token" form:"token"`
}
