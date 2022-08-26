package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/is0405/hacku/httputil"
	"github.com/is0405/hacku/util"
	"github.com/is0405/hacku/service"
	"github.com/is0405/hacku/model"
	"github.com/is0405/hacku/repository"
	"github.com/jmoiron/sqlx"
)

type Calender struct {
	db           *sqlx.DB
}

func NewCalender(db *sqlx.DB) *Calender {
	return &Calender{db: db}
}


// /carender/{recruitment_id} POST body : date time のリスト
func (a *Calender) PostRecCalender(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	rid, err := util.URLToInt(r)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	mcList := make([]model.Calender, 0);
	
	err = json.NewDecoder(r.Body).Decode(&mcList);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	recruitment, err := repository.GetRecruitmentFromRId(a.db, rid);
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if recruitment.SubmitId != getc.UserId {
		return http.StatusUnauthorized, nil, errors.New("you are not submitter")
	}

	CalenderService := service.NewCalender(a.db)
	for _, mc := range mcList {
		_, err = CalenderService.CreateCalender(&mc, rid);
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
	}

	return http.StatusOK, nil, nil
}

// /carender/{recruitment_id} GET
func (a *Calender) GetCalender(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	rid, err := util.URLToInt(r)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	calender , err := repository.GetCalenderFromRid(a.db, rid);
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var res []model.Calender
	for _, ins := range calender {
		num, err := repository.GetCountUserIdFromRidCid(a.db, rid, ins.Id)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		if num == 0 {
			ins.BookOk = true;
			ins.Iam = false;
		} else {
			uid, err := repository.GetUserIdFromRidCid(a.db, rid, ins.Id)
			if err != nil {
				return http.StatusInternalServerError, nil, err
			}

			ins.BookOk = false;
			ins.Iam = (uid == getc.UserId);
		}
		
		res = append(res, ins);
	}

	return http.StatusOK, res, nil 
}

// /carender/{recruitment_id} PATCH body:carender_id
func (a *Calender) PatchPartiCalender(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	rid, err := util.URLToInt(r)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	type reqCalender struct {
		id int `json:"id"`
	}
	var carenderId reqCalender
	err = json.NewDecoder(r.Body).Decode(&carenderId);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	CalenderService := service.NewCalender(a.db)
	_, err = CalenderService.PatchCalender(rid, getc.UserId, carenderId.id)
	
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	return http.StatusOK, nil, nil
}

// // /carender/{recruitment_id} DELETE
// func (a *Calender) DeletePartiCalender(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
// 	getc, err := httputil.GetClaimsFromContext(r.Context())
// 	if err != nil {
// 		return http.StatusInternalServerError, nil, err
// 	}

// 	rid, err := util.URLToInt(r)
// 	if err != nil {
// 		return http.StatusBadRequest, nil, err
// 	}

// 	var carenderId int
// 	err = json.NewDecoder(r.Body).Decode(&carenderId);
// 	if err != nil {
// 		return http.StatusBadRequest, nil, err
// 	}

// 	CalenderService := service.NewCalender(a.db)
// 	_, err = CalenderService.PatchCalender(rid, getc.UserId, carenderId)
// 	if err != nil {
// 		return http.StatusInternalServerError, nil, err
// 	}	
// 	return http.StatusOK, nil, nil
// }

