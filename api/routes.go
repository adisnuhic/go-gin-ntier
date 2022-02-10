package main

import (
	"github.com/adisnuhic/go-gin-ntier/controllers"
	middleware "github.com/adisnuhic/go-gin-ntier/middlewares"
	"github.com/golobby/container/pkg/container"
)

var (
	accountCtrl controllers.IAccountController
	healthCtrl  controllers.IHealthController
)

type Roles []int

var (
	RoleAdmin  = 1
	RoleWriter = 2
)

// initialize app routes
func initRoutes(c container.Container) {

	// Resolve dependencies and return concrete type of given abstractions
	c.Make(&accountCtrl)
	c.Make(&healthCtrl)

	// Group routes to specific version
	v1 := app.Group("/v1")

	// health check route
	healthRoutes := v1.Group("/ping")
	healthRoutes.GET("/", healthCtrl.Ping)

	// Account controller routes
	accountRoutes := v1.Group("/account")
	accountRoutes.POST("/register", accountCtrl.Register)
	accountRoutes.POST("/login", accountCtrl.Login)
	accountRoutes.POST("/refresh-token", accountCtrl.RefreshToken)

	protectedRoute := v1.Group("/me", middleware.Authorization(Roles{RoleAdmin, RoleWriter}...))
	protectedRoute.GET("/", accountCtrl.TestGetUserFromContext)

}
