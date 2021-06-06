package handler

type MailQueueRequest struct {
	To struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"to"`
	Cc []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"cc"`
	Bcc []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"bcc"`
	From struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	ReplyTo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"replyTo"`
	Content struct {
		Subject string `json:"subject"`
		Text    string `json:"text"`
	} `json:"content"`
}
