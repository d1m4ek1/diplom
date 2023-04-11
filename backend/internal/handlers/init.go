package handlers

import (
	"diplom/backend/middlewares"
	"html/template"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	csrf "github.com/utrack/gin-csrf"
)

type InitServer interface {
	InitServer() error
}

type files interface {
	initTemplates() (*template.Template, error)
	updateTemplates()
}

type routes interface {
	routesInit(db *sqlx.DB)

	routesInitAPI(db *sqlx.DB)
}

type controllers interface {
	home(db *sqlx.DB) gin.HandlerFunc

	admin() gin.HandlerFunc

	siteEdit() gin.HandlerFunc
}

type Server struct {
	Router         *gin.Engine
	ProductionMode bool
}

func (s *Server) InitServer(db *sqlx.DB) error {
	s.Router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	store, err := postgres.NewStore(db.DB, []byte("secret"))
	if err != nil {
		return err
	}

	go s.updateTemplates()

	s.Router.Static("/static", "./dist/static")

	s.Router.Use(sessions.Sessions("myssesion", store))

	s.Router.Use(csrf.Middleware(middlewares.CreateConfigCSRFDefenderMiddleware()))

	s.Router.Use(gin.Logger())
	s.Router.Use(gin.Recovery())

	s.routesInit(db)
	s.routesInitAPI(db)

	if err := s.Router.Run(":5050"); err != nil {
		return err
	}

	return nil
}
