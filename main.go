package main

import (
	"gin/routes"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	if utils.GetEnvData("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := routes.Routes()
	_ = r.Run(":9501")
}
