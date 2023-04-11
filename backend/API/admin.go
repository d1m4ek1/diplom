package api

import (
	"diplom/backend/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	csrf "github.com/utrack/gin-csrf"
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

func setSessionHasLoginToken(adminAuth *models.AuthAdmin, db *sqlx.DB, ctx *gin.Context) (bool, bool, int, error) {
	isAuth, status, err := adminAuth.LoginAuthAdmin(db)
	if err != nil {
		return false, false, status, err
	}

	if isAuth {
		session := sessions.Default(ctx)

		session.Options(sessions.Options{MaxAge: 60 * 60 * 24, HttpOnly: true, Secure: true, Path: "/", Domain: "localhost:3000"})
		session.Set("token", adminAuth.Token)

		if err := session.Save(); err != nil {
			return false, false, http.StatusInternalServerError, err
		}

		return true, true, http.StatusOK, err
	}

	return true, false, http.StatusOK, err
}

func AdminLogin(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var adminAuth models.AuthAdmin

		if err := ctx.ShouldBindJSON(&adminAuth); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"successfully": false,
				"error":        "Данные переданы не корректно",
			})
			return
		}

		successfully, isLogin, status, err := setSessionHasLoginToken(&adminAuth, db, ctx)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(status, gin.H{
				"successfully": false,
				"error":        "Неизвестная ошибка",
				"isLogin":      isLogin,
			})
			return
		}

		ctx.SetCookie("admin_token", adminAuth.Token, 60*60*24, "/", "localhost:3000", true, true)

		ctx.JSON(status, gin.H{
			"successfully": successfully,
			"isLogin":      isLogin,
		})
	})
}

func AdminLogout(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var adminAuth models.AuthAdmin
		var err error

		adminAuth.Token, err = ctx.Cookie("admin_token")
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		session := sessions.Default(ctx)

		session.Clear()

		if err := session.Save(); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		if status, err := adminAuth.LogoutAuthAdmin(db); err != nil {
			fmt.Println(err.Error())
			ctx.JSON(status, gin.H{
				"successfully": false,
				"error":        err.Error(),
			})
			return
		}

		ctx.SetCookie("admin_token", "", -1, "/", "localhost:3000", true, true)

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
		})
	})
}

func GetToken(db *sqlx.DB) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		ctx.Header("X-CSRF-TOKEN", csrf.GetToken(ctx))

		ctx.JSON(http.StatusOK, gin.H{
			"successfully": true,
		})
	})
}
