package repository

import (
	"database/sql"
	"github.com/is0405/hacku/model"
	"github.com/jmoiron/sqlx"
	//"fmt"
)

func GetUser(db *sqlx.DB, uid int) (*model.User, error) {
	var u model.User
	if err := db.Get(&u, `SELECT age, faculty, gender, mail, name FROM user WHERE id = ?;`, uid);
	err != nil {
		return nil, err
	}
	
	return &u, nil
}

func GetSubUser(db *sqlx.DB, suid int, code string) (*model.User, error) {
	var u model.User
	if err := db.Get(&u, `SELECT age, faculty, gender, mail, name,password FROM sub_user WHERE id = ? && code = ?;`, suid, code);
	err != nil {
		return nil, err
	}
	
	return &u, nil
}

func CreateSubUser(db *sqlx.DB, mu *model.User, code string) (sql.Result, error) {
	return db.Exec(`
INSERT INTO sub_user (age, faculty, gender, mail, name, password, code)
VALUES (?, ?, ?, ?, ?, ?, ?)
`, mu.Age, mu.Faculty, mu.Gender, mu.Mail, mu.Name, mu.Password, code)
}

func CreateUser(db *sqlx.DB, mu *model.User) (sql.Result, error) {
	return db.Exec(`
INSERT INTO user (age, faculty, gender, mail, name, password)
VALUES (?, ?, ?, ?, ?, ?)
`, mu.Age, mu.Faculty, mu.Gender, mu.Mail, mu.Name, mu.Password)
}

func RemoveUser(db *sqlx.DB, uid int) (sql.Result, error) {
	return db.Exec(`
DELETE FROM user WHERE id = ?;
`, uid)
}

func RemoveSubUser(db *sqlx.DB, uid int) (sql.Result, error) {
	return db.Exec(`
DELETE FROM sub_user WHERE id = ?;
`, uid)
}

func UpdateUser(db *sqlx.DB, mu *model.User) (sql.Result, error) {
	return db.Exec(`
UPDATE user SET age = ? WHERE id = ?;
`, mu.Age, mu.Id)
}

func UpdateUserPassword(db *sqlx.DB, mu *model.User) (sql.Result, error) {
	return db.Exec(`
UPDATE user SET password = ? WHERE id = ?;
`, mu.Password, mu.Id)
}


func GetParticipateUser(db *sqlx.DB, rid int, ruid int) ([]model.User, error) {
	a := make([]model.User, 0)
	if err := db.Select(&a, `SELECT age, faculty, gender, mail, name FROM user 
INNER JOIN participation ON user.id = participation.user_id
INNER JOIN recruitment ON participation.appeal_id = recruitment.id
WHERE recruitment.id = ? && recruitment.submit_id = ?;`, rid, ruid);
	err != nil {
		return nil, err
	}
	
	return a, nil
}
