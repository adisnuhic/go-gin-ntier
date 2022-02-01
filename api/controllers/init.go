package controllers

import (
	"github.com/adisnuhic/hearken/business"
	"github.com/golobby/container/pkg/container"
)

var (
	accountBl business.IAccountBusiness
)

// Bind controllers to IoC (dependency injection) container
func Init(c container.Container) {

	// Resolve dependencies and return concrete type of given abstractions
	c.Make(&accountBl)

	// Bind account controller
	c.Singleton(func() IAccountController {
		return NewAccountController(accountBl)
	})

}
