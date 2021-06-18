package qiscus_sdk

var (
	POST_COMMENT = "/api/v2.1/rest/post_comment"
)

type RequestWebhook struct {
}

type ParamQiscusSDK struct {
	AppID    string
	SecretID string
	BaseUrl  string
	BotID    string
}

type ParamSendMessage struct {
	UserID  string `form:"user_id" json:"user_id"`
	RoomID  string `form:"room_id" json:"room_id"`
	Message string `form:"message" json:"message"`
}

type ResponseSendMessage struct {
	Results struct {
		Comment struct {
			ID      int    `json:"id"`
			Message string `json:"message"`
		} `json:"comment"`
	} `json:"results"`

	Status int `json:"status"`
}
