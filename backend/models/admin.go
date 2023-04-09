package models

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Admin interface {
	GetAllQuestions() ([]FormData, int, error)
	DeleteQuestionByID(id int64) (int, error)
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

func (a *AdminData) DeleteQuestionByID(id int64) (int, error) {
	if _, err := a.DB.Exec(`
	DELETE FROM
		questions
	WHERE
		id=$1`, id); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
