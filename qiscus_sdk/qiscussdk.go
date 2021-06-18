package qiscus_sdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type QiscusSdkContract interface {
	SendMessage(request ParamSendMessage) (*ResponseSendMessage, error)
	GetBotID() string

	// Hit API SDK
	HitAPI(payload interface{}, url string, method string) ([]byte, error)
}

type qiscusSdk struct {
	ParamQiscusSDK
}

func NewQiscusSDK(request ParamQiscusSDK) QiscusSdkContract {
	return &qiscusSdk{
		request,
	}
}

func (sdk *qiscusSdk) SendMessage(request ParamSendMessage) (*ResponseSendMessage, error) {
	var response ResponseSendMessage

	data, err := sdk.HitAPI(request, POST_COMMENT, "POST")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (sdk *qiscusSdk) HitAPI(payload interface{}, url string, method string) ([]byte, error) {
	jsonPayload, _ := json.Marshal(&payload)

	reqhttp, err := http.NewRequest(method, sdk.BaseUrl+url, bytes.NewReader(jsonPayload))
	if err != nil {
		return nil, err
	}
	reqhttp.Header.Add("Content-Type", "application/json")
	reqhttp.Header.Add("QISCUS-SDK-APP-ID", sdk.AppID)
	reqhttp.Header.Add("QISCUS-SDK-SECRET", sdk.SecretID)

	res, err := http.DefaultClient.Do(reqhttp)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (sdk *qiscusSdk) GetBotID() string {
	if sdk.BotID == "" {
		sdk.BotID = "bot-qischat"
	}

	return sdk.BotID
}
