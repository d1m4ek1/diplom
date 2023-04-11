package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionID := session.Get("token")

		if sessionID == nil {
			ctx.Redirect(http.StatusSeeOther, "/admin/login")
			ctx.Abort()
		}
	})
}
