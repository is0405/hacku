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
)

type User struct {
	db           *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
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

func (a *User) CreateSubUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	mu := &model.User{}
	
	err := json.NewDecoder(r.Body).Decode(&mu);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if !util.CheckUser(mu) {
		return http.StatusUnprocessableEntity, nil, errors.New("required parameter is missing or invalid")
	}

	UserService := service.NewUser(a.db)
	_, err = UserService.CreateSubUser(mu)
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

	if !util.CheckUser(mu) {
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
