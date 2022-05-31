package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rafaelmbcosta/golan_gin_poc/controller"
	"github.com/rafaelmbcosta/golan_gin_poc/middlewares"
	"github.com/rafaelmbcosta/golan_gin_poc/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()

	setupLogOutput()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run()
}
