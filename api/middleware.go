package api

import (
	"errors"
	"net/http"
	"restapi/security"
	"strings"

	"github.com/gin-gonic/gin"
)

func authMiddleware(tokenBilder security.Builder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("authorization")
		if len(authHeader) == 0 {
			err := errors.New("falta token de autorización")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("formato de token inválido")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		if strings.ToLower(fields[0]) != "bearer" {
			err := errors.New("tipo de autorización no soportado:'bearer' requerido")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		accessToken := fields[1]
		payload, err := tokenBilder.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.Set("authorized", payload)
		ctx.Next()
	}
}
