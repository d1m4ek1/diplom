package models

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Filter interface {
	sortListHistory(db *sqlx.DB) ([]SiteContentItem, int, error)

	sortListQuestions(db *sqlx.DB) ([]FormData, int, error)

	SortList(db *sqlx.DB) (*SortedLists, int, error)
}

type SortListData struct {
	SortBy     string `json:"sortList"`
	SortByDate string `json:"sortListByDate"`
	NameList   string `json:"name_list"`
}

type SortedLists struct {
	SiteContentItems []SiteContentItem
	QuestionItems    []FormData
}

func generateQuery(nameList string) (sortby, sortbydate, allselect string) {
	var table string

	switch nameList {
	case "history":
		table = "site_content"
		break
	case "questions":
		table = "questions"
		break
	}

	sortby = fmt.Sprintf(`
	SELECT 
		* 
	FROM
		%s
	ORDER BY
		id
	DESC`, table)

	sortbydate = fmt.Sprintf(`
	SELECT 
		* 
	FROM
		%s
	ORDER BY
		add_date
	DESC`, table)

	allselect = fmt.Sprintf(`
	SELECT 
		* 
	FROM
		%s`, table)

	return sortby, sortbydate, allselect
}

func (s *SortListData) sortListHistory(db *sqlx.DB) ([]SiteContentItem, int, error) {
	var sortedLists []SiteContentItem
	sortby, sortbydate, allselect := generateQuery(s.NameList)

	switch "descending" {
	case s.SortBy:
		if err := db.Select(&sortedLists, sortby); err != nil {
			fmt.Println(err)
			return nil, http.StatusInternalServerError, err
		}

		break

	case s.SortByDate:
		if err := db.Select(&sortedLists, sortbydate); err != nil {
			fmt.Println(err)
			return nil, http.StatusInternalServerError, err
		}

		break
	}

	if s.SortBy == "ascending" && s.SortByDate == "ascending" {
		if err := db.Select(&sortedLists, allselect); err != nil {
			fmt.Println(err)
			return nil, http.StatusInternalServerError, err
		}
	}

	return sortedLists, http.StatusOK, nil
}

func (s *SortListData) sortListQuestions(db *sqlx.DB) ([]FormData, int, error) {
	var sortedLists []FormData
	sortby, sortbydate, allselect := generateQuery(s.NameList)

	switch "descending" {
	case s.SortBy:
		if err := db.Select(&sortedLists, sortby); err != nil {
			fmt.Println(err)
			return nil, http.StatusInternalServerError, err
		}

		break

	case s.SortByDate:
		if err := db.Select(&sortedLists, sortbydate); err != nil {
			fmt.Println(err)
			return nil, http.StatusInternalServerError, err
		}

		break
	}

	if s.SortBy == "ascending" && s.SortByDate == "ascending" {
		if err := db.Select(&sortedLists, allselect); err != nil {
			fmt.Println(err)
			return nil, http.StatusInternalServerError, err
		}
	}

	return sortedLists, http.StatusOK, nil
}

func (s *SortListData) SortList(db *sqlx.DB) (SortedLists, int, error) {
	var res SortedLists
	var err error
	var status int

	switch s.NameList {
	case "history":
		res.SiteContentItems, status, err = s.sortListHistory(db)
		if err != nil {
			return SortedLists{}, status, err
		}

		return res, status, nil

	case "questions":
		res.QuestionItems, status, err = s.sortListQuestions(db)
		if err != nil {
			return SortedLists{}, status, err
		}

		return res, http.StatusOK, nil
	}

	return SortedLists{}, http.StatusOK, nil
}
