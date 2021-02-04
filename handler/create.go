package handler

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type uuidPayload struct {
	UUID string `json:"uuid"`
}

// HandleCreate will handle a user joining a room
func (h *Handler) HandleCreate(c echo.Context) error {
	uuid := new(uuidPayload)
	if err := c.Bind(uuid); err != nil {
		return err
	}

	_, exists := rooms[uuid.UUID]
	if !exists {
		rooms[uuid.UUID] = make(map[*websocket.Conn]bool)
	}

	fmt.Println("uuid: ", uuid.UUID)
	return c.Redirect(301, "/")
}
