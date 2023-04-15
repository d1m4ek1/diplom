package handlers

import (
	api "diplom/backend/API"
	"diplom/backend/internal/middlewares"

	"github.com/jmoiron/sqlx"
)

func (s *Server) routesInit(db *sqlx.DB) {
	home := s.Router.Group("/")
	{
		home.GET("", s.home(db))
	}

	admin := s.Router.Group("/admin")
	{
		admin.GET("/login", s.adminLogin())

		admin.Use(middlewares.Auth()).GET("", s.admin())

		admin.Use(middlewares.Auth()).GET("site-edit", s.siteEdit())
	}
}

func (s *Server) routesInitAPI(db *sqlx.DB) {
	apix := s.Router.Group("/api")
	{
		form := apix.Group("/form")
		{
			form.POST("/send", api.SendForm(db))
		}

		admin := apix.Group("/admin")
		{
			auth := admin.Group("/auth")
			{
				auth.POST("/login", api.AdminLogin(db))

				auth.PATCH("/logout", api.AdminLogout(db))
			}

			admin.Use(middlewares.Auth()).GET("/all", api.GetAllQuestions(db))

			admin.Use(middlewares.Auth()).DELETE("/delete", api.DeleteQuestionByID(db))

			editor := admin.Group("/editor")
			editor.Use(middlewares.Auth())
			{
				editor.GET("/actual", api.GetActulSiteContent(db))

				editor.GET("/all", api.GetAllSiteContent(db))

				editor.GET("/sort_list", api.SortList(db))

				editor.PUT("/set_new", api.SetNewActualSiteContentFromHistory(db))

				editor.POST("/save", api.SaveContent(db))

				editor.DELETE("/delete", api.DeleteSiteContentFromHistory(db))
			}

			iframe := admin.Group("/iframe")
			iframe.Use(middlewares.Auth())
			{
				iframe.GET("", api.GetHTMLForIframe(db))
			}
		}

		token := apix.Group("/token")
		{
			token.GET("/get_token", api.GetToken(db))
		}
	}
}
