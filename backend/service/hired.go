package service

import (
	"github.com/is0405/hacku/dbutil"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	// "github.com/is0405/hacku/model"
	"github.com/is0405/hacku/repository"

	//"fmt"
)

type Hired struct {
	db *sqlx.DB
}

func NewHired (db *sqlx.DB) *Hired {
	return &Hired{db}
}

func (a *Hired) PostHired(aid int, uid int) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		hired, err := repository.PostHired(a.db, aid, uid)	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := hired.LastInsertId()
		if err != nil {
			return err
		}
		
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed auth insert transaction")
	}
	return createdId, nil
}

func (a *Hired) DeleteHired(aid int, uid int) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		hired, err := repository.DeleteHired(a.db, aid, uid)	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := hired.LastInsertId()
		if err != nil {
			return err
		}
		
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed auth insert transaction")
	}
	return createdId, nil
}
