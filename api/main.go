package main

import (
	"github.com/adisnuhic/hearken/business"
	"github.com/adisnuhic/hearken/config"
	"github.com/adisnuhic/hearken/controllers"
	"github.com/adisnuhic/hearken/db"
	"github.com/adisnuhic/hearken/initialize"
	"github.com/adisnuhic/hearken/repositories"
	"github.com/adisnuhic/hearken/services"
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
