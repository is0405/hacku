package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/is0405/hacku/httputil"
	"github.com/is0405/hacku/model"
	"github.com/is0405/hacku/service"
	"github.com/is0405/hacku/repository"
	"github.com/is0405/hacku/util"
	// "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"

	"fmt"
)

type User struct {
	db           *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

type SubUser struct {
	Id   int     `json:"id"`
	Code string  `json:"code"`
}

func (a *User) GetUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	res, err := repository.GetUser(a.db, getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	return http.StatusOK, res, nil
}

func (a *User) GetSubUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	su := &SubUser{}
	
	err := json.NewDecoder(r.Body).Decode(&su);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	// fmt.Println(su.Id)
	// fmt.Println(su.Code)
	mu, err := repository.GetSubUser(a.db, su.Id, su.Code)
	if err != nil {
		fmt.Println("err")
		return http.StatusInternalServerError, nil, err
	}
	
	return http.StatusOK, mu, nil
}

func (a *User) CreateSubUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	mu := &model.User{}
	
	err := json.NewDecoder(r.Body).Decode(&mu);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if !util.CheckUser(mu, true) {
		return http.StatusUnprocessableEntity, nil, errors.New("required parameter is missing or invalid")
	}

	mu.Password = util.HashGenerateSha256(mu.Password)
	code := util.CodeGenerate()

	cnt, err := repository.GetUserFromMail(a.db, mu.Mail)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if cnt != 0 {
		return http.StatusBadRequest, nil, errors.New("already existed")
	}

	UserService := service.NewUser(a.db)
	subId, err := UserService.CreateSubUser(mu, code)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}	

	res := SubUser {
		Id: int(subId),
		Code: code,	
	}
	
	return http.StatusOK, res, nil
}

func (a *User) CreateUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	su := &SubUser{}
	
	err := json.NewDecoder(r.Body).Decode(&su);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	mu, err := repository.GetSubUser(a.db, su.Id, su.Code)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	UserService := service.NewUser(a.db)
	_, err = UserService.CreateUser(su.Id, mu)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	return http.StatusOK, nil, nil
}

func (a *User) UpdateUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	mu := &model.User{}
	
	err = json.NewDecoder(r.Body).Decode(&mu);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if !util.CheckUser(mu, false) {
		return http.StatusUnprocessableEntity, nil, errors.New("required parameter is missing or invalid")
	}

	mu.Id = getc.UserId
	UserService := service.NewUser(a.db)
	_, err = UserService.UpdateUser(mu)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	res, err := repository.GetUser(a.db, mu.Id)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	return http.StatusOK, res, nil
}

func (a *User) DeleteUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	getc, err := httputil.GetClaimsFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	UserService := service.NewUser(a.db)
	_, err = UserService.DeleteUser(getc.UserId)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	return http.StatusOK, nil, nil
}

func (a *User) MailCheckUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	return http.StatusOK, nil, nil
}
