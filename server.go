package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelmbcosta/golan_gin_poc/controller"
	"github.com/rafaelmbcosta/golan_gin_poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run()
}
