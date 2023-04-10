package models

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Filter interface {
	SortList(db *sqlx.DB) ([]SiteContentItem, int, error)
}

type SortListData struct {
	SortBy     string `json:"sortList"`
	SortByDate string `json:"sortListByDate"`
}

func (s *SortListData) SortList(db *sqlx.DB) ([]SiteContentItem, int, error) {
	var items []SiteContentItem

	switch "descending" {

	case s.SortBy:
		if err := db.Select(&items, `
		SELECT 
			* 
		FROM
			site_content
		ORDER BY
			id
		DESC`); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		break

	case s.SortByDate:
		if err := db.Select(&items, `
		SELECT 
			* 
		FROM
			site_content
		ORDER BY
			add_date
		DESC`); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		break
	}

	if s.SortBy == "ascending" && s.SortByDate == "ascending" {
		if err := db.Select(&items, `
		SELECT 
			* 
		FROM
			site_content`); err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}

	return items, http.StatusOK, nil
}
