package middleware

import (
	// "errors"
	// "fmt"
	"fmt"
	// "log"
	"net/http"
	"strings"

	t "finance_tracker/auth_service/token"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		url := (ctx.Request.URL.Path)
		fmt.Println("token from heaf=der ", token)

		if strings.Contains(url, "swagger") || strings.Contains(url, "register") || strings.Contains(url, "login") || strings.Contains(url, "forgot_password") || strings.Contains(url, "resend") || strings.Contains(url, "confirm") {
			ctx.Next()
			return
		} else if isValid, err := t.ValidateToken(token); !isValid && err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}
