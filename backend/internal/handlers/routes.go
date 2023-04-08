package handlers

import "github.com/jmoiron/sqlx"

func (s *Server) routesInit() {
	home := s.Router.Group("/")
	{
		home.GET("", s.home())
	}
}

func (s *Server) routesInitAPI(db *sqlx.DB) {

}
