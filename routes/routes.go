package routes

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine) {
	r.GET("/ping", pingHandler)
}

func pingHandler(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "pong",
	})
}
