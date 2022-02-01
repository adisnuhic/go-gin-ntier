package services

import (
	"github.com/adisnuhic/hearken/config"
	"github.com/adisnuhic/hearken/repositories"
	"github.com/golobby/container/pkg/container"
)

var (
	userRepo         repositories.IUserRepository
	accountRepo      repositories.IAccountRepository
	authProviderRepo repositories.IAuthProviderRepository
	tokenRepo        repositories.ITokenRepository
)

// Bind services to IoC (dependency injection) container
func Init(c container.Container) {

	// Resolve dependencies and return concrete type of given abstractions
	c.Make(&userRepo)
	c.Make(&accountRepo)
	c.Make(&authProviderRepo)
	c.Make(&tokenRepo)

	// Bind user service
	c.Singleton(func() IUserService {
		return NewUserService(userRepo)
	})

	// Bind account service
	c.Singleton(func() IAccountService {
		return NewAccountService(accountRepo)
	})

	// Bind token service
	c.Singleton(func() ITokenService {
		return NewTokenService(tokenRepo)
	})

	// Bind auth provider service
	c.Singleton(func() IAuthProviderService {
		return NewAuthProviderService(authProviderRepo)
	})

	// Bind auth service
	c.Singleton(func() IAuthService {
		return NewAuthService(config.Load().JWTConf.Secret)
	})

}
