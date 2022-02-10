package main

import (
	"github.com/adisnuhic/go-gin-ntier/business"
	"github.com/adisnuhic/go-gin-ntier/config"
	"github.com/adisnuhic/go-gin-ntier/controllers"
	"github.com/adisnuhic/go-gin-ntier/db"
	"github.com/adisnuhic/go-gin-ntier/initialize"
	"github.com/adisnuhic/go-gin-ntier/repositories"
	"github.com/adisnuhic/go-gin-ntier/services"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container"
)

var app *gin.Engine

func main() {

	// Load conf
	cfg := config.Load()

	// Init db
	db.Init(cfg)

	// Create new DI container
	c := container.NewContainer()

	// Init repositories
	repositories.Init(c)

	// Init services
	services.Init(c)

	// Init businesses
	business.Init(c)

	// Init controllers
	controllers.Init(c)

	// Init GIN framework
	app = initialize.Gin()

	// Init routes
	initRoutes(c)

	// Run app
	app.Run(":8080")
}
