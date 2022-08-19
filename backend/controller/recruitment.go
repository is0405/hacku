package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	// "fmt"

	"github.com/is0405/hacku/httputil"
	"github.com/is0405/hacku/model"
	"github.com/is0405/hacku/util"
	"github.com/is0405/hacku/service"
	"github.com/is0405/hacku/repository"
	// "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

type Recruitment struct {
	db           *sqlx.DB
}

func NewRecruitment(db *sqlx.DB) *Recruitment {
	return &Recruitment{db: db}
}

type RecruitmentResponse struct {
	Recruitment      *model.Recruitment  `json:"recruitment"`
	NowParticipation int           `json:"nowparticipation"`
}

func (a *Recruitment) CreateRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	ma := &model.Recruitment{}
	
	err = json.NewDecoder(r.Body).Decode(&ma);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	
	ma.SubmitId = getc.UserId
	if ma.Conditions == "" {
		ma.Conditions = "特になし"
	}
	
	if !util.CheckRecruitment(ma) {
		return http.StatusUnprocessableEntity, nil, errors.New("required parameter is missing or invalid")
	}

	RecruitmentService := service.NewRecruitment(a.db)
	_, err = RecruitmentService.Create(ma)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}	

	
	return http.StatusOK, nil, nil
}

//特定の雇用情報をGet
func (a *Recruitment) GetRecruitmentFromID(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	_, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	aid, err := util.URLToInt(r)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	recruitment, err := repository.GetRecruitmentFromRId(a.db, aid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	cnt, err := repository.CountUidFromRid(a.db, aid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	res := RecruitmentResponse{
		Recruitment: recruitment,
		NowParticipation: cnt,
	}

	return http.StatusOK, res, nil
}

//お気に入りリストを用いて雇用リストを返す
func (a *Recruitment) GetRecruitmentListFromFavoriteList(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	ridList, err := repository.GetFavoriteRecruitementList(a.db, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	res := make([]RecruitmentResponse, 0)
	for _, rid := range ridList {
		recruitment, err := repository.GetRecruitmentFromRId(a.db, rid)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
		
		cnt, err := repository.CountUidFromRid(a.db, rid)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		ins := RecruitmentResponse{
			Recruitment: recruitment,
			NowParticipation: cnt,
		}

		res = append(res, ins)
	}
	
	return http.StatusOK, res, nil
}

//自分が提示した雇用情報を全て表示
func (a *Recruitment) GetMyAllRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	recruitments, err := repository.GetMyAllRecruitment(a.db, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var res []RecruitmentResponse
	for _, v := range recruitments {
		cnt, err := repository.CountUidFromRid(a.db, v.Id)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		ins := RecruitmentResponse{
			Recruitment: &v,
			NowParticipation: cnt,
		}
		
		res = append(res, ins)
	}

	return http.StatusOK, res, nil
}

func (a *Recruitment) GetOtherAllRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	recruitments, err := repository.GetMyAllRecruitment(a.db, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var res []RecruitmentResponse
	for _, v := range recruitments {
		cnt, err := repository.CountUidFromRid(a.db, v.Id)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		ins := RecruitmentResponse{
			Recruitment: &v,
			NowParticipation: cnt,
		}
		
		res = append(res, ins)
	}

	return http.StatusOK, res, nil
}

func (a *Recruitment) UpdateRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	ma := &model.Recruitment{}
	
	err = json.NewDecoder(r.Body).Decode(&ma);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if ma.SubmitId != getc.UserId {
		return http.StatusUnauthorized, nil, errors.New("Not Allowed Method")
	}
	
	id, err := util.URLToInt(r)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	
	mar, err := repository.GetRecruitmentFromRId(a.db, id)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	if ma.SubmitId != mar.SubmitId {
		return http.StatusUnauthorized, nil, errors.New("Not Allowed Method")
	}
	
	if !util.CheckRecruitment(ma) {
		return http.StatusUnprocessableEntity, nil, errors.New("required parameter is missing or invalid")
	}

	RecruitmentService := service.NewRecruitment(a.db)
	_, err = RecruitmentService.Update(ma)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}	

	res, err := repository.GetRecruitmentFromRId(a.db, ma.Id)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	return http.StatusOK, res, nil
}

func (a *Recruitment) DeleteRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	RecruitmentService := service.NewRecruitment(a.db)
	_, err = RecruitmentService.Delete(getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}	

	return http.StatusOK, nil, nil
}

func (a *Recruitment) GetParticipation(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	rid, err := util.URLToSecondInt(r);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	res, err := repository.GetParticipateUser(a.db, rid, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}
