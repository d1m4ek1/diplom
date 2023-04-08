package api

import (
	"diplom/backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SendForm(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var formData models.FormData

		if err := ctx.ShouldBindJSON(&formData); err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"successfully": false,
			})
			return
		}

		if status, err := formData.SetDataForm(db); err != nil {
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
