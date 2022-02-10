package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/adisnuhic/go-gin-ntier/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authorization middleware
func Authorization(roles ...int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := findAuthorizationToken(ctx.Request)
		var roleId int
		roleOk := false

		// Validate token
		if tokenStr != "" {
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.Load().JWTConf.Secret), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				roleId = int(claims["id"].(float64)) // here put roleId
			}

			// check roles
			if len(roles) > 0 {
				for _, r := range roles {
					if roleId == r {
						roleOk = true
						break
					}
				}
			}

			if len(roles) == 0 {
				roleOk = true
			}

			if err == nil && token.Valid && roleOk {
				ctx.Next()
			} else {
				ctx.AbortWithError(401, errors.New("unauthorized"))
				return
			}
		}

		if tokenStr == "" {
			ctx.AbortWithError(401, errors.New("unauthorized"))
			return
		}

	}
}

func findAuthorizationToken(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
