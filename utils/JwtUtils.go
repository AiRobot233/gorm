package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var h = 24 //默认的过期时间 24小时
var myKey = []byte("!Q@W#E$R5t")

type MyStandardClaims struct {
	Data                 any
	jwt.RegisteredClaims // v5版本新加的方法
}

// GetJwt 生成JWT
func GetJwt(data any, hour ...int) (bool, any) {
	if len(hour) > 0 {
		h = hour[0]
	}
	expireAt := time.Now().Add(time.Hour * time.Duration(h))
	claims := MyStandardClaims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),   // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString(myKey)
	if err != nil {
		return false, err.Error()
	}
	res := GetSlice()
	res["token"] = s
	res["expireAt"] = expireAt.Unix()
	res["user"] = data
	return true, res
}

// CheckJwt 解析JWT
func CheckJwt(tokenString string) (error, any) {
	t, err := jwt.ParseWithClaims(tokenString, &MyStandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if claims, ok := t.Claims.(*MyStandardClaims); ok && t.Valid {
		return nil, claims.Data
	} else {
		return err, nil
	}
}
