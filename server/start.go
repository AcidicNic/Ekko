package server

import (
	"log"
	"net/http"

	"github.com/AcidicNic/Ekko/handler"
)

// Start starts the server on  localhost:8000
func Start() {
	// create file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// configure websocket route
	http.HandleFunc("/ws", handler.HandleConnections)

	// start listening for messages
	go handler.HandleMessages()

	// start the server on port :8000
	log.Println("http server started on port :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
