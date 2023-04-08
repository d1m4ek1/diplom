package handlers

import (
	"diplom/backend/models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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
		ctx.HTML(http.StatusOK, "admin", gin.H{
			"IsQuestions":  true,
			"IsSiteEditor": false,
		})
	})
}

func (s *Server) siteEdit() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "site-edit", gin.H{
			"IsSiteEditor": true,
			"IsQuestions":  false,
		})
	})
}
