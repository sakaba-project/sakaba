package main

import (
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

func setupRouter(isRelease bool) *gin.Engine {
	if isRelease {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	return router
}

func main() {
	router := setupRouter(false)
	router.Run()
}
