package webhook

import (
	"github.com/kalfian/qisbot/app/webhook/handlers"
	"github.com/kalfian/qisbot/app/webhook/usecases"
	qiscussdk "github.com/kalfian/qisbot/qiscus_sdk"
)

func ProvideHandler(sdk qiscussdk.QiscusSdkContract) handlers.WebHookHandlerContract {
	usecase := usecases.NewWebHookUsecase(sdk)

	return handlers.NewWebhookHandler(usecase)
}
