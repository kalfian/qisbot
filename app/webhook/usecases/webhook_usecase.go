package usecases

import (
	"log"

	qiscussdk "github.com/kalfian/qisbot/qiscus_sdk"
)

type WebHookUsecaseContract interface {
	ResponseWebHook()
}

type webhookUsecase struct {
	SDK qiscussdk.QiscusSdkContract
}

func NewWebHookUsecase(sdk qiscussdk.QiscusSdkContract) WebHookUsecaseContract {
	return &webhookUsecase{
		SDK: sdk,
	}
}

func (usecase *webhookUsecase) ResponseWebHook() {
	log.Println("Webhook Hitten!")
}
