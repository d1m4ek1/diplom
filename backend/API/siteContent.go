package api

import (
	"diplom/backend/models"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetActulSiteContent(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var sti models.SiteContentItem

		if status, err := sti.GetActualSiteContent(db); err != nil {
			log.Println(err)
			ctx.JSON(status, gin.H{
				"successfully": false,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
			"data":         sti,
		})
	})
}

func GetAllSiteContent(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var stis models.SiteContentItems

		status, err := stis.GetAllSiteContent(db)
		if err != nil {
			log.Println(err)
			ctx.JSON(status, gin.H{
				"successfully": false,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
			"data":         stis.Items,
		})
	})
}

func SaveContent(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var sti models.SiteContentItem

		if err := ctx.ShouldBindJSON(&sti); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"successfully": false,
			})
			return
		}

		if status, err := sti.SaveContent(db); err != nil {
			log.Println(err)
			ctx.JSON(status, gin.H{
				"successfully": false,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
		})
	})
}

func SetNewActualSiteContentFromHistory(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var sti models.SiteContentItem

		if err := ctx.ShouldBindJSON(&sti); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"successfully": false,
			})
			return
		}

		if status, err := sti.SetNewActualSiteContentFromHistory(db); err != nil {
			log.Println(err)
			ctx.JSON(status, gin.H{
				"successfully": false,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
			"data":         sti,
		})
	})
}

func DeleteSiteContentFromHistory(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var sti models.SiteContentItem

		if err := ctx.ShouldBindJSON(&sti); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"successfully": false,
			})
			return
		}

		if status, err := sti.DeleteSiteContentFromHistory(db); err != nil {
			log.Println(err)
			ctx.JSON(status, gin.H{
				"successfully": false,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
		})
	})
}

func GetHTMLForIframe(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var siteContent models.SiteContentItem

		id, err := strconv.ParseInt(ctx.Query("idhtml"), 10, 32)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Problem", err)
			return
		}

		if status, err := siteContent.GetSiteContentByID(db, id); err != nil {
			ctx.String(status, "Problem", err)
			return
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"header":  siteContent.Header,
			"content": template.HTML(siteContent.Content),
		})
	})
}
