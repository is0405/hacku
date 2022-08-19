package repository

import (
	"database/sql"
	"github.com/is0405/hacku/model"
	"github.com/jmoiron/sqlx"
	//"fmt"
)

func CreateRecruitment(db *sqlx.DB, ma *model.Recruitment) (sql.Result, error) {
	return db.Exec(`
INSERT INTO recruitment (conditions, contents, max_participation, reward, submit_id, start_recruitment_period, finish_recruitment_period, start_implementation_period, finish_implementation_period, title, gender, min_age, max_age)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`, ma.Conditions, ma.Contents, ma.MaxParticipation, ma.Reward, ma.SubmitId, ma.StartRecruitmentPeriod, ma.FinishRecruitmentPeriod, ma.StartImplementationPeriod, ma.FinishImplementationPeriod, ma.Title, ma.Gender, ma.MinAge, ma.MaxAge)
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
UPDATE recruitment SET title = ?, contents = ?, conditions = ?, maxparticipation = ?, reward = ?, start_recruitment_period = ?, finish_recruitment_period = ?, start_implementation_period = ?, finish_implementation_period = ?, gender = ?, min_age = ?, max_age = ? WHERE id = ?;
`, ma.Title, ma.Contents, ma.Conditions, ma.MaxParticipation, ma.Reward, ma.StartRecruitmentPeriod, ma.FinishRecruitmentPeriod, ma.StartImplementationPeriod, ma.FinishImplementationPeriod, ma.Gender, ma.MinAge, ma.MaxAge, ma.Id)
}

