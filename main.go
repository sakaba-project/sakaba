package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"

	"github.com/sakaba-project/sakaba/config"
	"github.com/sakaba-project/sakaba/controllers"
	"github.com/sakaba-project/sakaba/db"
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

	u := router.Group("/users")
	{
		uc := controllers.NewUserController()
		u.GET("", uc.Index)
		u.POST("", uc.Create)
		u.GET("/:id", uc.Show)
		u.DELETE("/:id", uc.Delete)
	}

	return router
}

func main() {
	config.InitConfig()

	db.AutoMigrate()

	router := setupRouter(false)
	router.Run()
}
