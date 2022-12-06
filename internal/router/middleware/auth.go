package middleware

import (
	"github.com/gin-gonic/gin"
	"go-start/config"
	"go-start/internal/consts/e"
	"go-start/internal/pkg/app"
	"net/http"
)

func Auth(ctx *gin.Context) {
	token := ctx.GetHeader(config.Cfg.Jwt.TokenKey)
	if _, err := app.ParseUserByToken(token); err != nil {
		ctx.JSON(http.StatusUnauthorized, e.CodeMsg(e.UnAuthorized))
		ctx.Abort()
	}
	ctx.Next()
}
