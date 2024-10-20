package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(c *gin.Context) {})
		auth.POST("/register", func(c *gin.Context) {})
	}

	v2 := r.Group("/api")
	{
		v2.GET("/articles", func(c *gin.Context) {})
	}
	return r
}
