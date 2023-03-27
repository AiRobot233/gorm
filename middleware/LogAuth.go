package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type jsonResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w jsonResponseWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func (w jsonResponseWriter) Body() []byte {
	return w.body.Bytes()
}

func LogAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建自定义的 ResponseWriter 对象
		w := &jsonResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer([]byte{}),
		}
		// 替换 Gin 上下文的 Writer 字段为自定义的 ResponseWriter 对象
		c.Writer = w
		// 创建一个新的 Logrus 日志实例
		logger := logrus.New()
		// 设置日志输出格式
		logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.DateTime})
		file, _ := os.OpenFile(build(), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		// 将日志输出到文件中
		logger.SetOutput(file)
		// 记录请求信息
		logger.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"url":    c.Request.URL.String(),
			"ip":     c.ClientIP(),
		}).Info("request")
		// 记录响应信息
		c.Next()
		// 从自定义的 ResponseWriter 对象中读取响应体
		responseBody := w.Body()
		// 解析 JSON 数据
		var responseMap map[string]any
		err := json.Unmarshal(responseBody, &responseMap)
		if err != nil {
			return
		}
		if responseMap["error"].(float64) > 0 {
			logger.WithFields(logrus.Fields{
				"error":   1,
				"status":  c.Writer.Status(),
				"message": responseMap["message"],
			}).Info("response")
		}
	}
}

func build() string {
	day := time.Now().Format(time.DateOnly)
	filePath := "./log/"
	//判断文件夹是否存在
	_, err := os.Stat(filePath)
	if err != nil {
		_ = os.MkdirAll(filePath, os.ModePerm)
	}
	return filePath + day + ".log"
}
