package main

import (
	"github/beat-kuliah/sippad_backend/api"
)

func main() {
	server := api.NewServer(".")
	server.Start(8000)
}
