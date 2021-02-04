package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Server is a wrapper around our core
type Server struct {
	e       *echo.Echo
}

//New will create a new instance of the server.
func New() *Server {
	return &Server{
		e:       echo.New(), // new echo server to server the api
	}
}


// Start will start the server instance
func (s *Server) Start(port string) {
	// default port 8080
	if port == "" {
		port = ":8080"
	}

	// register routes
	s.Routes()

	// start server
	s.e.Logger.Fatal(s.e.Start(port))
}

// GetContext will return the context of the current echo server (mainly used for testing)
func (s *Server) GetContext(r *http.Request, w http.ResponseWriter) echo.Context {
	return s.e.NewContext(r, w)
}

// Stop will stop the server
func (s *Server) Stop() {
	// stop the server
	s.e.Close()
}