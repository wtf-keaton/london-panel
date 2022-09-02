package main

import (
	"template/internal/http_router"
	"template/internal/models"
	"template/pgk/memcache"
	"template/pgk/settings"
)

func init() {
	settings.Parse()
	models.Connect()

	go memcache.KeyCache.Fetch()
	go memcache.UserCache.Fetch()
	go memcache.CheatCache.Fetch()
	go memcache.BannedCache.Fetch()
}

func main() {
	if err := http_router.Serve().Listen(":4090"); err != nil {
		panic(err.Error())
	}
}
