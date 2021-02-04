package chat

// User object, holds information for a specific user
type User struct {
	Username string `json:"user"`
	Avatar   string `json:"avatar"`
}

// NewUser instatnitates a new user object
func NewUser(username, avatar string) User {
	return User{
		Username: username,
		Avatar:   avatar,
	}
}
