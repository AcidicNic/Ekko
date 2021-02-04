package main

import (
	"os"

	"github.com/AcidicNic/Ekko/server"
	"github.com/joho/godotenv"
)

func main() {
	// for development
	godotenv.Load()

	s := server.New()
	s.Start(os.Getenv("PORT"))
}
