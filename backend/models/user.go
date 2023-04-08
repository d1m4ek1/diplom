package models

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type User interface {
	SetDataForm(db *sqlx.DB) (int, error)
}

type FormData struct {
	Id         int64  `json:"_id" db:"id"`
	TypeUser   string `json:"typeUser" db:"type_user"`
	Question   string `json:"question" db:"question"`
	Email      string `json:"email" db:"email"`
	AddDate    string `json:"addDate" db:"add_date"`
	Tel        string `json:"tel" db:"tel"`
	FirstName  string `json:"firstName" db:"first_name"`
	SecondName string `json:"secondName" db:"second_name"`
	ThirdName  string `json:"thirdName" db:"third_name"`
}

func (f *FormData) SetDataForm(db *sqlx.DB) (int, error) {
	if err := db.Get(&f.Id, `
	INSERT INTO
		questions
		(type_user, question, email, add_date, tel, first_name)
	VALUES
		($1, $2, $3, $4, $5, $6) 
	RETURNING
		id`, f.TypeUser, f.Question, f.Email, f.AddDate, f.Tel, f.FirstName); err != nil {
		return http.StatusInternalServerError, err
	}

	if f.Id != 0 {
		if f.SecondName != "" {
			if _, err := db.Exec(`
			UPDATE
				questions
			SET
				second_name=$1
			WHERE
				id=$2`, f.SecondName, f.Id); err != nil {
				return http.StatusInternalServerError, err
			}
		}

		if f.ThirdName != "" {
			if _, err := db.Exec(`
			UPDATE
				questions
			SET
				third_name=$1
			WHERE
				id=$2`, f.ThirdName, f.Id); err != nil {
				return http.StatusInternalServerError, err
			}
		}
	}

	return http.StatusOK, nil
}
