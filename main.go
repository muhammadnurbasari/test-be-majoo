package main

import (
	"github.com/muhammadnurbasari/test-be-majoo/app/registry"
	"github.com/muhammadnurbasari/test-be-majoo/docs"
)

// !Important : Comments below are formatted as it is to be-
// read by Swagger tools (Swaggo)
// @title TEST BE MAJOO
// @version 0.0.1
// @description TEST BE MAJOO
// @BasePath
// @contact.name Muhammadnurbasari
// @contact.email m.nurbasari@gmail.com
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	appRegistry := registry.NewAppRegistry()
	appRegistry.StartServer()
}
