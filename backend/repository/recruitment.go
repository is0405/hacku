package repository

import (
	"database/sql"
	"github.com/is0405/hacku/model"
	"github.com/jmoiron/sqlx"
	//"fmt"
)

func CreateRecruitment(db *sqlx.DB, ma *model.Appeal) (sql.Result, error) {
	return db.Exec(`
INSERT INTO sub_user (conditions, contents, max_participation, reward, submit_id, term, title)
VALUES (?, ?, ?, ?, ?, ?, ?)
`, ma.Conditions, ma.Contents, ma.MaxParticipation, ma.Reward, ma.SubmitId, ma.Term, ma.Title)
}


func GetOtherAllRecruitment(db *sqlx.DB, myId int) ([]model.Appeal, error) {
	a := make([]model.Appeal, 0)
	if err := db.Select(&a, `
	SELECT * FROM recruitment WHERE submit_id != ?;
	`, myId); err != nil {
		return nil, err
	}
	
	return a, nil
}

func GetMyAllRecruitment(db *sqlx.DB, myId int) ([]model.Appeal, error) {
	a := make([]model.Appeal, 0)
	if err := db.Select(&a, `
	SELECT * FROM recruitment WHERE submit_id = ?;
	`, myId); err != nil {
		return nil, err
	}
	
	return a, nil
}

func GetRecruitmentFromRId(db *sqlx.DB, RId int) (*model.Appeal, error) {
	var a model.Appeal
	if err := db.Get(&a, `SELECT * FROM recruitment WHERE id = ?;
	`, RId); err != nil {
		return nil, err
	}
	
	return &a, nil
}

func RemoveRecruitment(db *sqlx.DB, rid int) (sql.Result, error) {
	return db.Exec(`
DELETE FROM appeal WHERE id = ?;
`, rid)
}

func UpdateRecruitment(db *sqlx.DB, ma *model.Appeal) (sql.Result, error) {
	return db.Exec(`
UPDATE appeal SET title = ?, contents = ?, conditions = ?, maxparticipation = ?, reward = ?, term = ? WHERE id = ?;
`, ma.Title, ma.Contents, ma.Conditions, ma.MaxParticipation, ma.Reward, ma.Term, ma.Id)
}

