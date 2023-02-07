package middleware

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"net/http"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var coder = base64.NewEncoding(base64Table)

func BaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//decryptPostFormParams(c)
		decryptPostJsonParams(c)
		c.Next()
	}
}

//解密post json参数
func decryptPostJsonParams(c *gin.Context) {
	j, err := getPostJsonParams(c)
	if err != nil {
		utils.Error(c, err.Error())
		c.Abort()
		return
	} else {
		if len(j) > 0 {
			for k, v := range j {
				d, baseErr := coder.DecodeString(v.(string))
				if baseErr != nil {
					utils.Error(c, baseErr.Error())
					c.Abort()
					return
				} else {
					j[k] = string(d)
				}
			}
			marshal, _ := json.Marshal(j) //转换为[]byte
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(marshal))
		}
	}
}

//解密post form参数
func decryptPostFormParams(c *gin.Context) {
	post, err := getPostFormParams(c)
	if err != nil {
		utils.Error(c, err.Error())
		c.Abort()
		return
	} else {
		if len(post) > 0 {
			for k, v := range post {
				d, baseErr := coder.DecodeString(v.(string))
				if baseErr != nil {
					utils.Error(c, baseErr.Error())
					c.Abort()
					return
				} else {
					c.Request.PostForm.Set(k, string(d))
				}
			}
		}
	}
}

//获取get参数
func getQueryParams(c *gin.Context) map[string]any {
	query := c.Request.URL.Query()
	var queryMap = make(map[string]any, len(query))
	for k := range query {
		queryMap[k] = c.Query(k)
	}
	return queryMap
}

//获取post form参数
func getPostFormParams(c *gin.Context) (map[string]any, error) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		if !errors.Is(err, http.ErrNotMultipart) {
			return nil, err
		}
	}
	var postMap = make(map[string]interface{}, len(c.Request.PostForm))
	for k, v := range c.Request.PostForm {
		if len(v) > 1 {
			postMap[k] = v
		} else if len(v) == 1 {
			postMap[k] = v[0]
		}
	}
	return postMap, nil
}

//获取post json参数
func getPostJsonParams(c *gin.Context) (map[string]any, error) {
	params := utils.GetSlice()
	if err := c.ShouldBindBodyWith(&params, binding.JSON); err == nil {
		return params, nil
	} else {
		return nil, err
	}
}
