package controller

import (
	// "encoding/json"
	// "errors"
	"net/http"

	// "github.com/is0405/hacku/httputil"
	"github.com/is0405/hacku/model"
	// "github.com/is0405/hacku/repository"
	// "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

type Recruitment struct {
	db           *sqlx.DB
}

func NewRecruitment(db *sqlx.DB) *Recruitment {
	return &Recruitment{db: db}
}

func (a *Recruitment) CreateRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusOK, nil, nil
}

func (a *Recruitment) GetRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	//response構造体に用意したデータを格納
	res := &model.Appeal{
		Conditions: "aa",
		Contents: "ss",
		CreatedAt: "2022-2-10",
		Id: 1,
		MaxParticipate: 2,
		Reward:"時給1円",
		SubmitId: 1,
		Term: "2022-1-10~2022-10-24",
		Title:"N実験",
		UpdatedAt:"2022-2-21",
	}

	return http.StatusOK, res, nil
}

func (a *Recruitment) UpdateRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	//response構造体に用意したデータを格納
	res := &model.Appeal{
		Conditions: "aa",
		Contents: "ss",
		CreatedAt: "2022-2-10",
		Id: 1,
		MaxParticipate: 2,
		Reward:"時給1円",
		SubmitId: 1,
		Term: "2022-1-10~2022-10-24",
		Title:"N実験",
		UpdatedAt:"2022-2-21",
	}

	return http.StatusOK, res, nil
}

func (a *Recruitment) DeleteRecruitment(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	return http.StatusOK, nil, nil
}

func (a *Recruitment) GetParticipation(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	//response構造体に用意したデータを格納
	var res []model.User;
	instance := model.User{
		Age: 20,
		CreatedAt: "2022-1-10",
		Faculty: 1,
		Gender: 1,
		Id: 1,
		Mail: "is0000@aaa.com",
		Name: "NN",
		UpdatedAt: "2022-2-10",
	}

	res = append(res, instance)
	
	return http.StatusOK, res, nil
}
