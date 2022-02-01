package main

import (
	"github.com/adisnuhic/hearken/controllers"
	middleware "github.com/adisnuhic/hearken/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/pkg/container"
)

var (
	accountCtrl controllers.IAccountController
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

	v1 := app.Group("/v1")

	// Account controller routes
	accountRoutes := v1.Group("/account")
	accountRoutes.POST("/register", accountCtrl.Register)
	accountRoutes.POST("/login", accountCtrl.Login)
	accountRoutes.POST("/refresh-token", accountCtrl.RefreshToken)

	protectedRoute := v1.Group("/me", middleware.Authorization(Roles{RoleAdmin, RoleWriter}...))
	protectedRoute.GET("/test", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

}
