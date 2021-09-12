package jwt

import (
	"errors"
	"time"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/dgrijalva/jwt-go"
)

var ErrorInvalidToekn = errors.New("invalid token")

const TokenExpireDuration = time.Hour * 24 * 7

var jwtKey = []byte("dojnb")

// MyClaims 自定义声明结构体并内嵌 jwt.StandardClaims
// jwt 包

type MyClaims struct {
	StudentID string `json:"student_id"`
	jwt.StandardClaims
}

// 发放 token
func ReleaseToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(TokenExpireDuration)
	claims := &MyClaims{
		StudentID: user.StudentID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "DOJ",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// 解析 token
func ParseToekn(tokenString string) (*MyClaims, error) {
	var claims = new(MyClaims)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return claims, nil
	}

	return nil, ErrorInvalidToekn
}
