package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LogAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求开始时间
		startTime := time.Now()
		c.Next()
		// 请求结束时间
		endTime := time.Now()
		// 请求响应时间
		latencyTime := endTime.Sub(startTime)
		go func() {
			// 创建一个新的 Logrus 日志实例
			logger := logrus.New()
			// 设置日志输出格式
			logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.DateTime})
			file, _ := os.OpenFile(build(), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
			// 将日志输出到文件中
			logger.SetOutput(file)
			logger.WithFields(logrus.Fields{
				"client_ip":  c.ClientIP(),
				"method":     c.Request.Method,
				"status":     c.Writer.Status(),
				"latency":    latencyTime.Milliseconds(),
				"user_agent": c.Request.UserAgent(),
				"path":       c.Request.URL.String(),
			}).Info("ask")
		}()
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
