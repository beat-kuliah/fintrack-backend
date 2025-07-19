package main

import (
	"github/beat-kuliah/fintrack_backend/api"
)

func main() {
	server := api.NewServer(".")
	server.Start(8000)
}
