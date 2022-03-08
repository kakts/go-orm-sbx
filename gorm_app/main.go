package main

import (
	"github.com/kakts/go-orm-sbx/gorm_app/repository"
	"github.com/kakts/go-orm-sbx/gorm_app/server"
)

// startup server with initializing db
func main() {
	repository.InitDB()
	server.Init()
}
