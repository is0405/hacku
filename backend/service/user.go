package service

import (
	"github.com/is0405/hacku/dbutil"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/is0405/hacku/model"
	"github.com/is0405/hacku/repository"

	//"fmt"
)

type User struct {
	db *sqlx.DB
}

func NewUser (db *sqlx.DB) *User {
	return &User{db}
}

func (a *User) CreateSubUser(smu *model.User, code string) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		user, err := repository.CreateSubUser(a.db, smu, code)	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := user.LastInsertId()
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


func (a *User) CreateUser(sid int, mu *model.User) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
	
		user, err := repository.CreateUser(a.db, mu)	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := user.LastInsertId()
		if err != nil {
			return err
		}
		
		createdId = id

		_, err = repository.RemoveSubUser(a.db, sid)
		if err != nil {
			return err
		}
		
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed auth insert transaction")
	}
	return createdId, nil
}

func (a *User) UpdateUser(mu *model.User) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
	
		user, err := repository.UpdateUser(a.db, mu)	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := user.LastInsertId()
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

func (a *User) DeleteUser(sid int) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
	
		_, err := repository.RemoveSubUser(a.db, sid)
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		createdId = 0
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed auth insert transaction")
	}
	return createdId, nil
}
