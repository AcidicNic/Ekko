package server

import (
	"path/filepath"

	"github.com/AcidicNic/Ekko/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Routes handles all of the routes for the application
func (s *Server) Routes() {
	// CORS
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
	}))
	r := handler.NewHandler()

	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Static(filepath.Join("./public")))

	s.e.GET("/ws", r.HandleConnections)

	s.e.POST("/room/exists", r.RoomExists)
	s.e.POST("/create/room", r.HandleCreate)

	go r.HandleMessages()
}
