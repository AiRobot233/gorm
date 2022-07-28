package middleware

import (
	"bytes"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func MostBaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params map[string]interface{}     //声明变量，不分配内存
		params = make(map[string]interface{}) //必可不少，分配内存
		if err := c.ShouldBind(&params); err == nil {
			d, _ := coder.DecodeString(params["encrypt"].(string))
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(d))
		} else {
			utils.Error(c, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
