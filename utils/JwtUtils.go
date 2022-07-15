package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var h = 24 //默认的过期时间 24小时
var myKey = []byte("!Q@W#E$R5t")

type MyStandardClaims struct {
	Data interface{}
	jwt.StandardClaims
}

//获取jwt
func GetJwt(data interface{}, hour ...int) string {
	if len(hour) > 0 {
		h = hour[0]
	}
	ms := MyStandardClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(h)).Unix(),
			Issuer:    "hc",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &ms)
	signedString, err := token.SignedString(myKey)
	if err != nil {
		return err.Error()
	}
	return signedString
}

//解析验证jwt
func CheckJwt(signedString string) (bool, interface{}) {
	claims, err := jwt.ParseWithClaims(signedString, &MyStandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return true, err.Error()
	}
	return false, claims.Claims.(*MyStandardClaims).Data
}
