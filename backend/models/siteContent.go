package models

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type SiteContent interface {
	GetActualSiteContent(db *sqlx.DB) (int, error)

	SaveContent(db *sqlx.DB) (int, error)

	GetAllSiteContent(db *sqlx.DB) (int, error)
}

type SiteContentItem struct {
	Id      int64  `json:"_id" db:"id"`
	Header  string `json:"header" db:"header"`
	Content string `json:"content" db:"content"`
	AddDate string `json:"addDate" db:"add_date"`
}

type SiteContentItems struct {
	Items []SiteContentItem
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

func (s *SiteContentItems) GetAllSiteContent(db *sqlx.DB) (int, error) {
	if err := db.Select(&s.Items, `
	SELECT 
		*
  FROM 
		site_content`); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusInternalServerError, nil
}

func (s *SiteContentItem) SaveContent(db *sqlx.DB) (int, error) {
	if _, err := db.Exec(`
	INSERT INTO
		site_content
		(header, content, add_date)
	VALUES
		($1, $2, $3)`, s.Header, s.Content, s.AddDate); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
