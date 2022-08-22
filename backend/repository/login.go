package repository

import (
	"github.com/is0405/hacku/model"
	"github.com/jmoiron/sqlx"
)

func CheckLogin(db *sqlx.DB, lb *model.LoginBody) (int, error) {
	var a int
	if err := db.Get(&a, `
SELECT id FROM user WHERE mail = ? AND password = ?
`, lb.Mail, lb.Password); err != nil {
		return 0, err
	}
	return a, nil
}
