package handlers

import (
	"github.com/gin-gonic/gin"
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

func (uc *webhookHandler) GetClientSendMessage(c *gin.Context) {

}
