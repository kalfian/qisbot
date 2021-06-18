package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/kalfian/qisbot/app/webhook/models"
	qiscussdk "github.com/kalfian/qisbot/qiscus_sdk"
)

type WebHookUsecaseContract interface {
	ResponseWebHook(request models.RequestWebHook) error
}

type webhookUsecase struct {
	SDK qiscussdk.QiscusSdkContract
}

func NewWebHookUsecase(sdk qiscussdk.QiscusSdkContract) WebHookUsecaseContract {
	return &webhookUsecase{
		SDK: sdk,
	}
}

func (usecase *webhookUsecase) ResponseWebHook(request models.RequestWebHook) error {
	// SEND Message to client
	var paramBotChat qiscussdk.ParamSendMessage
	paramBotChat.UserID = usecase.SDK.GetBotID()
	paramBotChat.RoomID = strconv.Itoa(request.Payload.Room.ID)

	if request.Payload.Message.Text == "hai" {
		paramBotChat.Message = "Haloo!"
		_, err := usecase.SDK.SendMessage(paramBotChat)

		if err != nil {
			return err
		}
	} else if request.Payload.Message.Text == "hitung" {
		for i := 0; i < 10; i++ {
			paramBotChat.Message = fmt.Sprintf("AYOO %d", i+1)
			_, err := usecase.SDK.SendMessage(paramBotChat)

			if err != nil {
				return err
			}
		}
	} else {
		paramBotChat.Message = "Apa?"
		_, err := usecase.SDK.SendMessage(paramBotChat)

		if err != nil {
			return err
		}
	}

	payload, _ := json.Marshal(&request)
	log.Printf("Webhook payload: %s", payload)

	return nil
}
