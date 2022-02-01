package repositories

import (
	"github.com/adisnuhic/hearken/db"
	"github.com/golobby/container/pkg/container"
)

// Bind repositories to IoC (dependency injection) container
func Init(c container.Container) {

	// Bind user repository
	c.Singleton(func() IUserRepository {
		return NewUserRepository(db.Connection())
	})

	// Bind account repository
	c.Singleton(func() IAccountRepository {
		return NewAccountRepository(db.Connection())
	})

	// Bind token repository
	c.Singleton(func() ITokenRepository {
		return NewTokenRepository(db.Connection())
	})

	// Bind auth provider repository
	c.Singleton(func() IAuthProviderRepository {
		return NewAuthProviderRepository(db.Connection())
	})
}
