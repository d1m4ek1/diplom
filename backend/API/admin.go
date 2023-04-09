package api

import (
	"diplom/backend/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetAllQuestions(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		admin := models.AdminData{DB: db}

		if status, err := admin.GetAllQuestions(); err != nil {
			log.Println(err)
			ctx.JSON(status, gin.H{
				"successfully": false,
				"items":        nil,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
			"items":        admin.QuestionItems,
		})
	})
}

func DeleteQuestionByID(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		admin := models.AdminData{DB: db}

		id, err := strconv.ParseInt(ctx.Query("id"), 10, 32)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"successfully": false,
			})
			return
		}

		if status, err := admin.DeleteQuestionByID(id); err != nil {
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
