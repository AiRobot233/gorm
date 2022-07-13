package controller

import (
	"context"
	"fmt"
	"gin/config"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(context *gin.Context) {

}

func Redis() {
	rdb := config.BuildRedis()
	ctx := context.Background()
	sth, _ := time.ParseDuration("1h")
	res := rdb.Set(ctx, "key", "asdadada", sth)
	fmt.Println(res)
}
