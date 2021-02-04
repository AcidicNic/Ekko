package handler

import (
	"net/http"

	"github.com/AcidicNic/Ekko/chat"
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

	_, exists := h.Rooms[uuid.UUID]
	if !exists {
		h.Rooms[uuid.UUID] = chat.NewRoom(uuid.UUID)
		return c.Redirect(301, "/")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/create")
}
