package repository

import (
	"database/sql"
	// "github.com/is0405/hacku/model"
	"github.com/jmoiron/sqlx"
)

func CountFromRidUid(db *sqlx.DB, rid int, uid int) (int, error) {
	var a int
	if err := db.Get(&a, `
	SELECT COUNT(id) FROM participation WHERE user_id = ? && recruitment_id = ?;
	`, uid, rid); err != nil {
		return 0, err
	}
	
	return a, nil
}

func PostHired(db *sqlx.DB, aid int, uid int) (sql.Result, error) {
	return db.Exec(`
INSERT INTO participation (recruitment_id, user_id)
VALUES (?, ?)
`, aid, uid)
}

func DeleteHired(db *sqlx.DB, aid int, uid int) (sql.Result, error) {
	return db.Exec(`
DELETE FROM participation WHERE user_id = ? && recruitment_id = ?;
`, uid, aid)
}

func CountUidFromRid(db *sqlx.DB, aid int) (int, error) {
	var a int
	if err := db.Get(&a, `
	SELECT COUNT(user_id) FROM participation WHERE recruitment_id = ?;
	`, aid); err != nil {
		return 0, err
	}
	
	return a, nil
}


func GetAllRidFromUid(db *sqlx.DB, uid int) ([]int, error) {
	a := make([]int, 0)
	if err := db.Select(&a, `
	SELECT recruitment_id FROM participation WHERE user_id = ?;
	`, uid); err != nil {
		return nil, err
	}
	
	return a, nil
}
