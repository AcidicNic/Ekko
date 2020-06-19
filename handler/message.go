package handler

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Avatar   string `json:"avatar"`
}
