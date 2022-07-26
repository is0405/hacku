package service

import (
	"github.com/is0405/hacku/dbutil"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/is0405/hacku/model"
	"github.com/is0405/hacku/repository"

	//"fmt"
)

type Recruitment struct {
	db *sqlx.DB
}

func NewRecruitment (db *sqlx.DB) *Recruitment {
	return &Recruitment{db}
}

func (a *Recruitment) Create(ma *model.ReqRecruitment) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		rec, err := repository.CreateRecruitment(a.db, ma)	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := rec.LastInsertId()
		if err != nil {
			return err
		}

		rid := int(id)

		for _, dl:= range ma.DateList {
			dateId, cnt, err := repository.GetDateIdAndCount(a.db, dl.Date, dl.Time);	
			if err != nil {
				return err
			}

			if cnt == 0 {
				d, err := repository.CreateDate(a.db, dl.Date, dl.Time);
				if err != nil {
					return err
				}
			
				ccid, err := d.LastInsertId()
				if err != nil {
					return err
				}
				dateId = int(ccid)
			}

			_, err = repository.CreateCalender(a.db, dateId, rid);	
			if err != nil {
				return err
			}
		}
		
		
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed auth insert transaction")
	}
	return createdId, nil
}

func (a *Recruitment) Update(ma *model.Recruitment) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		rec, err := repository.UpdateRecruitment(a.db, ma)	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := rec.LastInsertId()
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

func (a *Recruitment) Delete(rid int) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		rma, err := repository.RemoveRecruitment(a.db, rid)	
		if err != nil {
			return err
		}
		
		if err := tx.Commit(); err != nil {
			return err
		}
		
		id, err := rma.LastInsertId()
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
