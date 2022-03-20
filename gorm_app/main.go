package main

import (
	_ "github.com/kakts/go-orm-sbx/gorm_app/docs"
	"github.com/kakts/go-orm-sbx/gorm_app/repository"
	"github.com/kakts/go-orm-sbx/gorm_app/server"
)

// @title REST API with gorm
// @version 1.0
// @description This is a sandbox project for REST API written in golang.
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /
// startup server with initializing db
func main() {
	// テストとわけるため、repositoryはインターフェースにして分ける
	repository.InitDB()
	server.Init()
}
