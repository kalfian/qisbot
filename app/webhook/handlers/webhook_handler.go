package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kalfian/qisbot/app/webhook/models"
	"github.com/kalfian/qisbot/app/webhook/usecases"
)

type WebHookHandlerContract interface {
	GetClientSendMessage(c *gin.Context)
}

type webhookHandler struct {
	Uc usecases.WebHookUsecaseContract
}

func NewWebhookHandler(uc usecases.WebHookUsecaseContract) WebHookHandlerContract {
	return &webhookHandler{
		Uc: uc,
	}
}

func (handler *webhookHandler) GetClientSendMessage(c *gin.Context) {
	var request models.RequestWebHook

	if err := c.Bind(&request); err != nil {
		log.Printf("Error Bind request %v \n", err)
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Error Get Request",
		})
	}

	if err := handler.Uc.ResponseWebHook(request); err != nil {
		log.Printf("Error Response webhook %v \n", err)
		c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": "Error Response webhook",
		})
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Webhook success!",
	})
}
