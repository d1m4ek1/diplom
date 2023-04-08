package models

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type SiteContent interface {
	GetActualSiteContent(db *sqlx.DB) (int, error)
}

type SiteContentItem struct {
	Id      int64  `json:"_id" db:"id"`
	Header  string `json:"header" db:"header"`
	Content string `json:"content" db:"content"`
	AddDate string `json:"addDate" db:"add_date"`
}

func (s *SiteContentItem) GetActualSiteContent(db *sqlx.DB) (int, error) {
	if err := db.Get(s, `
	SELECT 
		*
  FROM 
		site_content
  ORDER BY 
		add_date DESC
  LIMIT 
		1`); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
