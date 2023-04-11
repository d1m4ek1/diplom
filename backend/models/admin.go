package models

import (
	generatetoken "diplom/backend/pkg/generateToken"
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

type AuthAdmin struct {
	AdminID  int64  `json:"-" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
	Token    string `json:"-" db:"token"`
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

func (a *AuthAdmin) LoginAuthAdmin(db *sqlx.DB) (bool, int, error) {
	var isAuth bool
	var err error

	a.Token, err = generatetoken.CreateToken(256)
	if err != nil {
		return false, http.StatusInternalServerError, err
	}

	if err := db.Get(&isAuth, `SELECT EXISTS (SELECT * FROM admin WHERE password=$1 AND login=$2)`, a.Password, a.Login); err != nil {
		return false, http.StatusInternalServerError, err
	}

	if err := db.Get(&a.AdminID, `SELECT id FROM admin WHERE password=$1 AND login=$2`, a.Password, a.Login); err != nil {
		return false, http.StatusInternalServerError, err
	}

	if _, err := db.Exec(`
			UPDATE
				admin
			SET
				token=$1
			WHERE
				id=$2`, a.Token, a.AdminID); err != nil {
		return false, http.StatusInternalServerError, err
	}

	return isAuth, http.StatusOK, nil
}

func (a *AuthAdmin) LogoutAuthAdmin(db *sqlx.DB) (int, error) {
	if _, err := db.Exec("UPDATE admin SET token='' WHERE token=$1", a.Token); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
