package main

import (
	"github.com/kakts/go-orm-sbx/gorm/repository"
	"github.com/kakts/go-orm-sbx/gorm/server"
)

// startup server with initializing db
func main() {
	// テストとわけるため、repositoryはインターフェースにして分ける
	repository.InitDB()
	server.Init()
}
