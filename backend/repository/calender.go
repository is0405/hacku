package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/is0405/hacku/model"
	//"fmt"
)

func CreateCalender(db *sqlx.DB, dateId int, rid int) (sql.Result, error) {
	return db.Exec(`
INSERT INTO calender (date_id, recruitment_id)
VALUES (?, ?)
`, dateId, rid);
}

func CreateDate(db *sqlx.DB, date string, time string) (sql.Result, error) {
	return db.Exec(`
INSERT INTO calender_date (date, time)
VALUES (?, ?)
`, date, time);
}

func GetDateIdAndCount(db *sqlx.DB, date string, time string) (int, int, error) {
	var a int
	if err := db.Get(&a, `SELECT COUNT(id) FROM calender_date WHERE date = ? AND time = ?;
	`, date, time); err != nil {
		return -1, 0, err
	}

	if a != 0 {
		if err := db.Get(&a, `SELECT id FROM calender_date WHERE date = ? AND time = ?;
	`, date, time); err != nil {
			return -1, 1, err
		}
		return a, 1, nil
	}
	
	return -1, 0, nil
}

func GetCalenderFromRid(db *sqlx.DB, rid int) ([]model.Calender, error) {
	a := make([]model.Calender, 0);
	if err := db.Select(&a, `
SELECT calender.id, calender_date.date, calender_date.time
FROM calender 
JOIN calender_date ON calender_date.id = calender.date_id
WHERE recruitment_id = ?;
	`, rid); err != nil {
		return nil, err
	}
	return a, nil
}

func UpdateParticipation(db *sqlx.DB, rid int, userId int, carenderId int) (sql.Result, error) {
	return db.Exec(`
UPDATE participation SET calender_id = ? WHERE recruitment_id = ? AND user_id = ?;
`, carenderId, rid, userId)
}

func GetUserIdFromRidCid(db *sqlx.DB, rid int, cid int) (int, error) {
	var a int
	if err := db.Get(&a, `
SELECT user_id
FROM participation 
WHERE recruitment_id = ? and calender_id = ?;
	`, rid, cid); err != nil {
		return 0, err
	}
	return a, nil
}

func GetCountUserIdFromRidCid(db *sqlx.DB, rid int, cid int) (int, error) {
	var a int
	if err := db.Get(&a, `
SELECT COUNT(user_id)
FROM participation 
WHERE recruitment_id = ? and calender_id = ?;
	`, rid, cid); err != nil {
		return 0, err
	}
	return a, nil
}
