package api

import (
	"diplom/backend/models"
	"log"
	"net/http"

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
