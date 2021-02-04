package main

import "github.com/AcidicNic/Ekko/server"

func main() {
	s := server.New()
	s.Start(":8080")
}
