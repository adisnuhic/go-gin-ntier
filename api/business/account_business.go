package business

import (
	"time"

	"github.com/adisnuhic/go-gin-ntier/ecode"
	"github.com/adisnuhic/go-gin-ntier/models"
	apperror "github.com/adisnuhic/go-gin-ntier/pkg"
	"github.com/adisnuhic/go-gin-ntier/services"
)

const (
	// AuthProviderLocal local provider
	AuthProviderLocal = "local"

	// RefreshTokenTypeID refresh token ID
	RefreshTokenTypeID = 1
)

// IAccountBusiness interface
type IAccountBusiness interface {
	Register(firstName string, lastName string, email string, password string) (string, string, *apperror.AppError)
	Login(email string, password string) (*models.User, string, string, *apperror.AppError)
	RefreshToken(refreshToken string) (*models.User, string, string, *apperror.AppError)
	AuthenticateUser(user *models.User) (string, string, *apperror.AppError)
}

type accountBusiness struct {
	Service             services.IAccountService
	UserService         services.IUserService
	AuthProviderService services.IAuthProviderService
	AuthService         services.IAuthService
	TokenService        services.ITokenService
}

// NewAccountBusiness -
func NewAccountBusiness(svc services.IAccountService, userSvc services.IUserService, authProviderSvc services.IAuthProviderService, authSvc services.IAuthService, tokenSvc services.ITokenService) IAccountBusiness {
	return &accountBusiness{
		Service:             svc,
		UserService:         userSvc,
		AuthProviderService: authProviderSvc,
		AuthService:         authSvc,
		TokenService:        tokenSvc,
	}
}

// AuthenticateUser authenticates user
func (bl accountBusiness) AuthenticateUser(user *models.User) (string, string, *apperror.AppError) {
	// access token
	accessToken, errToken := bl.AuthService.GenerateAccessToken(user.ID, user.Email)
	if errToken != nil {
		return "", "", errToken
	}

	// refresh token
	refreshToken, errRefreshToken := bl.TokenService.CreateRefreshToken(user.ID, user.Email)
	if errRefreshToken != nil {
		return "", "", errRefreshToken
	}

	return accessToken, refreshToken.Token, nil
}

// Register user
func (bl accountBusiness) Register(firstName string, lastName string, email string, password string) (string, string, *apperror.AppError) {
	exists := bl.UserService.Exists(email)
	if exists {
		return "", "", apperror.New(ecode.ErrUserExistsCode, ecode.ErrUserExistsMsg, ecode.ErrUserExistsMsg)
	}

	// Generate hash
	hash, errHash := bl.AuthService.GeneratePasswordHash(password)
	if errHash != nil {
		return "", "", errHash
	}

	// insert user data into db
	user := &models.User{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		IsConfirmed: false,
	}

	user, errUser := bl.Service.Register(user)
	if errUser != nil {
		return "", "", errUser
	}

	// insert auth data into db
	auth := &models.AuthProvider{
		Provider: AuthProviderLocal,
		UserID:   user.ID,
		UID:      hash,
	}

	errAuth := bl.AuthProviderService.Save(auth)
	if errAuth != nil {
		return "", "", errAuth
	}

	// authenticate user
	accessToken, refreshToken, err := bl.AuthenticateUser(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Login user
func (bl accountBusiness) Login(email string, password string) (*models.User, string, string, *apperror.AppError) {
	user, errUser := bl.UserService.GetByEmail(email)

	if errUser != nil {
		return nil, "", "", apperror.New(ecode.ErrLoginFailedCode, errUser.Cause, ecode.ErrLoginFailedMsg)
	}

	// get user data
	auth, errAuth := bl.AuthProviderService.GetByUserID(user.ID)
	if errAuth != nil {
		return nil, "", "", errAuth
	}

	// check if hash matches
	if ok := bl.AuthService.ComparePasswordHash(password, auth.UID); ok == false {
		return nil, "", "", apperror.New(ecode.ErrLoginFailedCode, errAuth.Cause, ecode.ErrLoginFailedMsg)
	}

	// authenticate user
	accessToken, refreshToken, err := bl.AuthenticateUser(user)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil

}

// RefreshToken
func (bl accountBusiness) RefreshToken(refreshToken string) (*models.User, string, string, *apperror.AppError) {
	token, err := bl.TokenService.GetByToken(refreshToken)
	if err != nil {
		return nil, "", "", err
	}

	user, err := bl.UserService.GetByID(token.UserID)
	if err != nil {
		return nil, "", "", apperror.New(ecode.ErrUnableToFetchUserCode, err.Cause, ecode.ErrUnableToFetchUserMsg)
	}

	// check if token is refresh token
	if token.TokenType != RefreshTokenTypeID {
		return nil, "", "", apperror.New(ecode.ErrNotRefreshTokenCode, err.Cause, ecode.ErrNotRefreshTokenMsg)
	}

	// check if token has expired
	if token.ExpiresAt.Before(time.Now().UTC()) {
		return nil, "", "", apperror.New(ecode.ErrTokenExpiredCode, err.Cause, ecode.ErrTokenExpiredMsg)
	}

	// authenticate user
	accessToken, refreshToken, err := bl.AuthenticateUser(user)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}
