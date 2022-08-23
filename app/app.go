package main

import (
	"template/internal/http_router"
	"template/internal/models"
	"template/pgk/settings"
)

func init() {
	settings.Parse()
	models.Connect()
}

func main() {
	if err := http_router.Serve().Listen(":4090"); err != nil {
		panic(err.Error())
	}
}
