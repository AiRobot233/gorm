package middleware

import (
	"bytes"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"io"
)

func MostBaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := utils.GetSlice()
		if err := c.ShouldBind(&params); err == nil {
			d, _ := coder.DecodeString(params["encrypt"].(string))
			c.Request.Body = io.NopCloser(bytes.NewBuffer(d))
		} else {
			utils.Error(c, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
