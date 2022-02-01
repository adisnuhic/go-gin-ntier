package controllers

import (
	"net/http"

	"github.com/adisnuhic/hearken/viewmodels"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// BaseController -
type BaseController struct {
}

// Render -
func (ctrl BaseController) Render(ctx *gin.Context, success bool, status int, data interface{}, err interface{}) {
	ctx.Status(status)

	ctx.JSON(status, viewmodels.Response{
		Success:   success,
		RequestID: requestid.Get(ctx),
		Data:      data,
		Error:     err,
	})
}

// RenderSuccess renders success response
func (ctrl BaseController) RenderSuccess(ctx *gin.Context, data interface{}) {
	ctrl.Render(ctx, true, http.StatusOK, data, nil)
}

// RenderError renders error response
func (ctrl BaseController) RenderError(ctx *gin.Context, err interface{}) {
	ctrl.Render(ctx, false, http.StatusInternalServerError, nil, err)
}

// RednerBadRequest renders bad request response
func (ctrl BaseController) RednerBadRequest(ctx *gin.Context, err interface{}) {
	ctrl.Render(ctx, false, http.StatusBadRequest, nil, err)
}
