package controllers

import (
	"github.com/adisnuhic/go-gin-ntier/viewmodels"
	"github.com/gin-gonic/gin"
)

// IHealthController interface
type IHealthController interface {
	Ping(ctx *gin.Context)
}

type healthController struct {
	BaseController
}

// NewHealthController -
func NewHealthController() IHealthController {
	return &healthController{}
}

// Ping returns PONG health check data
func (ctrl healthController) Ping(ctx *gin.Context) {
	ctrl.RenderSuccess(ctx, viewmodels.Response{
		Success: true,
		Data:    "PONG",
	})
}
