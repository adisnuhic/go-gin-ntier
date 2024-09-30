package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/adisnuhic/go-gin-ntier/config"
	"github.com/adisnuhic/go-gin-ntier/models"
	apperror "github.com/adisnuhic/go-gin-ntier/pkg"
	"github.com/adisnuhic/go-gin-ntier/viewmodels"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

// GetUserFromContext returns logged in user from context
func (ctrl BaseController) GetUserFromContext(ctx *gin.Context) (*models.User, *apperror.AppError) {
	user := &models.User{}
	bearer := ctx.Request.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		bearer = bearer[7:]
	}

	if bearer == "" {
		return nil, apperror.New(401, "unauthorized", "unauthorized")
	}

	if bearer != "" {
		token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Load().JWTConf.Secret), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user.ID = uint64(claims["id"].(float64))
			user.Email = fmt.Sprintf("%v", claims["email"])
		}

		if err == nil && token.Valid {
			return user, nil
		}
	}

	return nil, apperror.New(401, "unauthorized", "unauthorized")
}
