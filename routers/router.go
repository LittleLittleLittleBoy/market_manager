package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.StaticFS("/html", http.Dir("static"))
	r.StaticFS("/images", http.Dir("images"))

	apiGroup := r.Group("api")
	{
		apiGroup.POST("/upload", uploadImage)
		apiGroup.POST("/identityAccount", identityAccount)
	}

	return r
}
