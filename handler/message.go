package handler

// Message object, defines attributes that a message will have
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Avatar   string `json:"avatar"`
}
