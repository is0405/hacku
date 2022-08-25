package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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
	Id               int    `json:"recruitmentId"`
	Name             string `json:"name"`
	Faculty          int    `json:"faculty"`
	UpadateAt        string `json:"date"`
	Title            string `json:"title"`
	Content          string `json:"content"`
	Conditions       string `json:"conditions"`
	MaxSubjects      int    `json:"maxSubjects"`
	Period           string `json:"period"`
	Reward           string `json:"reward"`
	Sex              int    `json:"sex"`
	MinAge           int    `json:"minAge"`
	MaxAge           int    `json:"maxAge"`
	NowParticipation int    `json:"nowSubjects"`
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

	rid, err := util.URLToInt(r)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	recruitment, err := repository.GetRecruitmentFromRId(a.db, rid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	//何人参加しているか
	cnt, err := repository.CountUidFromRid(a.db, rid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	//雇用主の情報をとってくる
	user, err := repository.GetUser(a.db, recruitment.SubmitId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	res := RecruitmentResponse{
		Id: recruitment.Id,
		Name: user.Name,
	    Faculty: user.Faculty, 
	    UpadateAt: recruitment.UpdatedAt,
	    Title: recruitment.Title,
	    Content: recruitment.Contents,
		Conditions: recruitment.Conditions,
	    MaxSubjects: recruitment.MaxParticipation,
	    Period: recruitment.Period,
	    Reward: recruitment.Reward,
	    Sex: recruitment.Gender,
	    MinAge: recruitment.MinAge,
	    MaxAge: recruitment.MaxAge,
	    NowParticipation:cnt,
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

	user, err := repository.GetUser(a.db, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var res []RecruitmentResponse
	for _, recruitment := range recruitments {
		cnt, err := repository.CountUidFromRid(a.db, recruitment.Id)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		ins := RecruitmentResponse{
			Id: recruitment.Id,
			Name: user.Name,
			Faculty: user.Faculty, 
			UpadateAt: recruitment.UpdatedAt,
			Title: recruitment.Title,
			Content: recruitment.Contents,
			Conditions: recruitment.Conditions,
			MaxSubjects: recruitment.MaxParticipation,
			Period: recruitment.Period,
			Reward: recruitment.Reward,
			Sex: recruitment.Gender,
			MinAge: recruitment.MinAge,
			MaxAge: recruitment.MaxAge,
			NowParticipation:cnt,
		}
		
		res = append(res, ins)
	}

	return http.StatusOK, res, nil
}

//自分以外の雇用情報表示
func (a *Recruitment) GetOtherAllRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	query := r.URL.Query();
	gender, err := strconv.Atoi(query["sex"][0]);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	
	minAge, err := strconv.Atoi(query["minAge"][0]);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	
	maxAge, err := strconv.Atoi(query["maxAge"][0]);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	
	recruitments, err := repository.GetOtherAllRecruitment(a.db, getc.UserId, gender, minAge, maxAge)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var res []RecruitmentResponse
	for _, recruitment := range recruitments {
		cnt, err := repository.CountUidFromRid(a.db, recruitment.Id)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		//雇用主の情報をとってくる
		user, err := repository.GetUser(a.db, recruitment.SubmitId)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
	
		ins := RecruitmentResponse{
			Id: recruitment.Id,
			Name: user.Name,
			Faculty: user.Faculty, 
			UpadateAt: recruitment.UpdatedAt,
			Title: recruitment.Title,
			Content: recruitment.Contents,
			Conditions: recruitment.Conditions,
			MaxSubjects: recruitment.MaxParticipation,
			Period: recruitment.Period,
			Reward: recruitment.Reward,
			Sex: recruitment.Gender,
			MinAge: recruitment.MinAge,
			MaxAge: recruitment.MaxAge,
			NowParticipation:cnt,
		}
		
		res = append(res, ins)
	}

	return http.StatusOK, res, nil
}

//自分が参加する雇用情報を全て表示
func (a *Recruitment) GetMyParticipationAllRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	rid_list, err := repository.GetAllRidFromUid(a.db, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}	

	var res []RecruitmentResponse
	for _, rid := range rid_list {
		recruitment, err := repository.GetRecruitmentFromRId(a.db, rid)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
		
		cnt, err := repository.CountUidFromRid(a.db, recruitment.Id)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		//雇用主の情報をとってくる
		user, err := repository.GetUser(a.db, recruitment.SubmitId)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
	
		ins := RecruitmentResponse{
			Id: recruitment.Id,
			Name: user.Name,
			Faculty: user.Faculty, 
			UpadateAt: recruitment.UpdatedAt,
			Title: recruitment.Title,
			Content: recruitment.Contents,
			Conditions: recruitment.Conditions,
			MaxSubjects: recruitment.MaxParticipation,
			Period: recruitment.Period,
			Reward: recruitment.Reward,
			Sex: recruitment.Gender,
			MinAge: recruitment.MinAge,
			MaxAge: recruitment.MaxAge,
			NowParticipation:cnt,
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

//雇用者が参加者の情報をとってくる
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
