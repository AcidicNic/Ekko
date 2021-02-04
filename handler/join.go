package handler

import (
	"github.com/labstack/echo/v4"
)

// RoomExists handler function to return a bool if the room exists or not
func (h *Handler) RoomExists(c echo.Context) error {
	uuid := new(uuidPayload)
	if err := c.Bind(uuid); err != nil {
		return err
	}

	_, exists := h.Rooms[uuid.UUID]
	if !exists {
		return c.JSON(200, false)
	}

	return c.JSON(200, true)
}
