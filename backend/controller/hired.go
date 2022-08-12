package controller

import (
	// "encoding/json"
	"errors"
	"net/http"

	"github.com/is0405/hacku/httputil"
	"github.com/is0405/hacku/util"
	"github.com/is0405/hacku/service"
	// "github.com/is0405/hacku/model"
	"github.com/is0405/hacku/repository"
	// "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

type Hired struct {
	db           *sqlx.DB
}

func NewHired(db *sqlx.DB) *Hired {
	return &Hired{db: db}
}


func (a *Hired) PostHired(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	aid, err := util.URLToInt(r)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	
	mr, err := repository.GetRecruitmentFromRId(a.db, aid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	//自分が登録していない
	cnt, err := repository.CountFromRidUid(a.db, aid, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if cnt == 0 {
		return http.StatusBadRequest, nil, errors.New("already applied")
	}

	nowParticipation, err := repository.CountUidFromRid(a.db, aid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if mr.MaxParticipation <= nowParticipation {
		return http.StatusBadRequest, nil, errors.New("max_participation")
	}

	HiredService := service.NewHired(a.db)
	_, err = HiredService.PostHired(aid, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}

func (a *Hired) DeleteHired(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	aid, err := util.URLToInt(r)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	cnt, err := repository.CountFromRidUid(a.db, aid, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if cnt != 1 {
		return http.StatusBadRequest, nil, errors.New("already applied")
	}

	HiredService := service.NewHired(a.db)
	_, err = HiredService.DeleteHired(aid, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}

