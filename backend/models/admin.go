package models

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Admin interface {
	GetAllQuestions() ([]FormData, int, error)
}

type AdminData struct {
	DB            *sqlx.DB
	QuestionItems []FormData
}

func (a *AdminData) GetAllQuestions() (int, error) {
	if err := a.DB.Select(&a.QuestionItems, `
	SELECT
		*
	FROM
		questions`); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
