package repository

import (
	"database/sql"
	"github.com/is0405/hacku/model"
	"github.com/jmoiron/sqlx"
	//"fmt"
)

func CreateRecruitment(db *sqlx.DB, ma *model.Recruitment) (sql.Result, error) {
	return db.Exec(`
INSERT INTO recruitment (conditions, contents, max_participation, reward, submit_id, period, title, min_age, max_age, gender)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`, ma.Conditions, ma.Contents, ma.MaxParticipation, ma.Reward, ma.SubmitId, ma.Period, ma.Title, ma.MinAge, ma.MaxAge, ma.Gender)
}


func GetOtherAllRecruitment(db *sqlx.DB, myId int) ([]model.Recruitment, error) {
	
	a := make([]model.Recruitment, 0)
	if err := db.Select(&a, `
	SELECT * FROM recruitment WHERE submit_id != ?;
	`, myId); err != nil {
		return nil, err
	}
	
	return a, nil
}

func GetMyAllRecruitment(db *sqlx.DB, myId int) ([]model.Recruitment, error) {
	a := make([]model.Recruitment, 0)
	if err := db.Select(&a, `
	SELECT * FROM recruitment WHERE submit_id = ?;
	`, myId); err != nil {
		return nil, err
	}
	
	return a, nil
}

func GetRecruitmentFromRId(db *sqlx.DB, RId int) (*model.Recruitment, error) {
	var a model.Recruitment
	if err := db.Get(&a, `SELECT * FROM recruitment WHERE id = ?;
	`, RId); err != nil {
		return nil, err
	}
	
	return &a, nil
}

func RemoveRecruitment(db *sqlx.DB, rid int) (sql.Result, error) {
	return db.Exec(`
DELETE FROM recruitment WHERE id = ?;
`, rid)
}

func UpdateRecruitment(db *sqlx.DB, ma *model.Recruitment) (sql.Result, error) {
	return db.Exec(`
UPDATE recruitment SET title = ?, contents = ?, conditions = ?, maxparticipation = ?, reward = ?, period = ?, min_age, max_age, gender WHERE id = ?;
`, ma.Title, ma.Contents, ma.Conditions, ma.MaxParticipation, ma.Reward, ma.Period, ma.Id, ma.MinAge, ma.MaxAge, ma.Gender)
}

