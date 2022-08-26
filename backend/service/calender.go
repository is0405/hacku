package service

import (
	"github.com/is0405/hacku/dbutil"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/is0405/hacku/model"
	"github.com/is0405/hacku/repository"

	//"fmt"
)

type Calender struct {
	db *sqlx.DB
}

func NewCalender (db *sqlx.DB) *Calender {
	return &Calender{db}
}

func (a *Calender) CreateCalender(mc *model.Calender, rid int) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		dateId, cnt, err := repository.GetDateIdAndCount(a.db, mc.Date, mc.Time);	
		if err != nil {
			return err
		}

		if cnt == 0 {
			d, err := repository.CreateCalender(a.db, dateId, rid);
			if err != nil {
				return err
			}
			
			id, err := d.LastInsertId()
			if err != nil {
				return err
			}
			dateId = int(id)
		}

		c, err := repository.CreateCalender(a.db, dateId, rid);	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := c.LastInsertId()
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

func (a *Calender) PatchCalender(rid int, userId int, carenderId int) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateParticipation(a.db, rid, userId, carenderId);	
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
