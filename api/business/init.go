package business

import (
	"github.com/adisnuhic/hearken/services"
	"github.com/golobby/container/pkg/container"
)

var (
	userSvc         services.IUserService
	accountSvc      services.IAccountService
	authSvc         services.IAuthService
	authProviderSvc services.IAuthProviderService
	tokenSvc        services.ITokenService
)

// Bind business to IoC (dependency injection) container
func Init(c container.Container) {

	// Resolve dependencies and return concrete type of given abstractions
	c.Make(&userSvc)
	c.Make(&accountSvc)
	c.Make(&authSvc)
	c.Make(&authProviderSvc)
	c.Make(&tokenSvc)

	// Bind user business
	c.Singleton(func() IUserBusiness {
		return NewUserBusiness(userSvc)
	})

	// Bind account business
	c.Singleton(func() IAccountBusiness {
		return NewAccountBusiness(accountSvc, userSvc, authProviderSvc, authSvc, tokenSvc)
	})

	// Bind user business
	c.Singleton(func() IUserBusiness {
		return NewUserBusiness(userSvc)
	})

}
