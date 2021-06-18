package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kalfian/qisbot/app/webhook"
	qiscussdk "github.com/kalfian/qisbot/qiscus_sdk"
)

func InitRoute(r *gin.Engine, sdk qiscussdk.QiscusSdkContract) {
	webHookProvider := webhook.ProvideHandler(sdk)

	r.GET("/ping", pingHandler)

	// WEB HOOK
	r.POST("/webhook", webHookProvider.GetClientSendMessage)
}

func pingHandler(c *gin.Context) {
	c.JSON(200, map[string]string{
		"message": "pong",
	})
}
