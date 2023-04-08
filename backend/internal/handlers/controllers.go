package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) home() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
