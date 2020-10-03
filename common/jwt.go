package common

import (
	"ginessential/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_key")

//Claims  json web token claim
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

// ReleaseToken 发送token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gin-vue-demo",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// TODO jwk生成的字符串包含三部分,由.分隔开
// TODO eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjQsImV4cCI6MTYwMjMyNjY3OCwiaWF0IjoxNjAxNzIxODc4LCJpc3MiOiJnaW4tdnVlLWRlbW8iLCJzdWIiOiJ1c2VyIHRva2VuIn0.jNiAzyK01XHhIozwtM8O0ErZnr--35EdsKP76Y8QYwI
// TODO echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9  | base64 -d                 -----> 			{"alg":"HS256","typ":"JWT"}
// TODO echo 第二段 | base64 -d  ------> 		{"UserID":4,"exp":1602326678,"iat":1601721878,"iss":"gin-vue-demo","sub":"user token"}
// TODO
