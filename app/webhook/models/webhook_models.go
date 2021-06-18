package models

/*
* This request use to grab response from
* qiscus api webhoook
 */
type RequestWebHook struct {
	Type    string `form:"type"`
	Payload struct {
		From struct {
			ID    int    `form:"id"`
			Email string `form:"email"`
			Name  string `form:"name"`
		} `form:"from"`

		Room struct {
			ID      int    `form:"id"`
			TopicID int    `form:"topic_id"`
			Type    string `form:"type"`
			Name    string `form:"name"`
		} `form:"from"`

		Message struct {
			Type string `form:"type"`
			Text string `form:"text"`
		} `form:"room"`
	} `form:"message"`
}
