package controllers

import (
	"github.com/adisnuhic/hearken/business"
	"github.com/adisnuhic/hearken/requests"
	"github.com/adisnuhic/hearken/viewmodels"
	"github.com/gin-gonic/gin"
)

// IAccountController interface
type IAccountController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}

type accountController struct {
	BaseController
	Business business.IAccountBusiness
}

// NewAccountController -
func NewAccountController(business business.IAccountBusiness) IAccountController {
	return &accountController{
		Business: business,
	}
}

// Register user
func (ctrl accountController) Register(ctx *gin.Context) {
	reqObj := requests.Registration{}

	if err := ctx.ShouldBindJSON(&reqObj); err != nil {
		ctrl.RednerBadRequest(ctx, err)
		return
	}

	accessToken, refreshToken, appErr := ctrl.Business.Register(reqObj.FirstName, reqObj.LastName, reqObj.Email, reqObj.Password)
	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		return
	}

	ctrl.RenderSuccess(ctx, viewmodels.Auth{
		User:         nil,
		Token:        accessToken,
		RefreshToken: refreshToken,
	})
}

// Login user
func (ctrl accountController) Login(ctx *gin.Context) {
	reqObj := requests.Login{}

	if err := ctx.ShouldBindJSON(&reqObj); err != nil {
		ctrl.RednerBadRequest(ctx, err.Error())
		ctx.Abort()
		return
	}

	user, accessToken, refreshToken, appErr := ctrl.Business.Login(reqObj.Email, reqObj.Password)

	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		ctx.Abort()
		return
	}

	ctrl.RenderSuccess(ctx, &viewmodels.Auth{
		User:         user,
		Token:        accessToken,
		RefreshToken: refreshToken,
	})

}

// RefreshToken -
func (ctrl accountController) RefreshToken(ctx *gin.Context) {
	reqObj := requests.RefreshToken{}

	if err := ctx.ShouldBindJSON(&reqObj); err != nil {
		ctrl.RednerBadRequest(ctx, err.Error())
		return
	}

	user, accessToken, refreshToken, appErr := ctrl.Business.RefreshToken(reqObj.Token)
	if appErr != nil {
		ctrl.RenderError(ctx, appErr)
		return
	}

	ctrl.RenderSuccess(ctx, &viewmodels.Auth{
		User:         user,
		Token:        accessToken,
		RefreshToken: refreshToken,
	})

}
