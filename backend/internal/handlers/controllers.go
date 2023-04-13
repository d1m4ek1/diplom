package handlers

import (
	"diplom/backend/models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	csrf "github.com/utrack/gin-csrf"
)

func (s *Server) home(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var siteContent models.SiteContentItem

		if status, err := siteContent.GetActualSiteContent(db); err != nil {
			ctx.String(status, "Problem", err)
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"header":  siteContent.Header,
			"content": template.HTML(siteContent.Content),
		})
	})
}

func (s *Server) admin() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin", nil)
	})
}

func (s *Server) siteEdit() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin", nil)
	})
}

func (s *Server) adminLogin() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin_login", gin.H{
			"csrf": csrf.GetToken(ctx),
		})
	})
}

func (s *Server) settings() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin", nil)
	})
}
