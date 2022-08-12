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
func GetJwt(data interface{}, hour ...int) (bool, interface{}) {
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
	var res map[string]interface{}     //声明变量，不分配内存
	res = make(map[string]interface{}) //必可不少，分配内存
	res["token"] = signedString
	res["expireAt"] = expireAt
	res["user"] = data
	return true, res
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
