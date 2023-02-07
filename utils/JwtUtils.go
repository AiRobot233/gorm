package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var h = 24 //默认的过期时间 24小时
var myKey = []byte("!Q@W#E$R5t")

type MyStandardClaims struct {
	Data any
	jwt.StandardClaims
}

// GetJwt 获取jwt
func GetJwt(data any, hour ...int) (bool, any) {
	if len(hour) > 0 {
		h = hour[0]
	}
	expireAt := time.Now().Add(time.Hour * time.Duration(h)).Unix()
	ms := MyStandardClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    "hc",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &ms)
	signedString, err := token.SignedString(myKey)
	if err != nil {
		return false, err
	}
	res := GetSlice()
	res["token"] = signedString
	res["expireAt"] = expireAt
	res["user"] = data
	return true, res
}

// CheckJwt 解析验证jwt
func CheckJwt(signedString string) (bool, any) {
	claims, err := jwt.ParseWithClaims(signedString, &MyStandardClaims{}, func(token *jwt.Token) (any, error) {
		return myKey, nil
	})
	if err != nil {
		return true, err.Error()
	}
	return false, claims.Claims.(*MyStandardClaims).Data
}
