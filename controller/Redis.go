package controller

import (
	"context"
	"fmt"
	"gin/config"
	"time"
)

func Redis() {
	rdb := config.BuildRedis()
	ctx := context.Background()
	sth, _ := time.ParseDuration("1h")
	res := rdb.Set(ctx, "key", "asdadada", sth)
	fmt.Println(res)
}
