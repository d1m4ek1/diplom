package handlers

import (
	api "diplom/backend/API"

	"github.com/jmoiron/sqlx"
)

func (s *Server) routesInit(db *sqlx.DB) {
	home := s.Router.Group("/")
	{
		home.GET("", s.home(db))
	}

	admin := s.Router.Group("/admin")
	{
		admin.GET("", s.admin())

		admin.GET("site-edit", s.siteEdit())
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
			admin.GET("/all", api.GetAllQuestions(db))

			admin.DELETE("/delete", api.DeleteQuestionByID(db))

			editor := admin.Group("/editor")
			{
				editor.GET("/actual", api.GetActulSiteContent(db))

				editor.GET("/all", api.GetAllSiteContent(db))

				editor.PUT("/set_new", api.SetNewActualSiteContentFromHistory(db))

				editor.POST("/save", api.SaveContent(db))

				editor.DELETE("/delete", api.DeleteSiteContentFromHistory(db))
			}

			iframe := s.Router.Group("/iframe")
			{
				iframe.GET("", api.GetHTMLForIframe(db))
			}
		}
	}
}
