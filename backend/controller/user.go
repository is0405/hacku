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

type User struct {
	db           *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}


func (a *User) GetUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	//response構造体に用意したデータを格納
	res := &model.User{
		Age: 20,
		CreatedAt: "2022-1-10",
		Faculty: 1,
		Gender: 1,
		Id: 1,
		Mail: "is0000@aaa.com",
		Name: "NN",
		UpdatedAt: "2022-2-10",
	}
	
	return http.StatusOK, res, nil
}

func (a *User) CreateUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	return http.StatusOK, nil, nil
}

func (a *User) UpdateUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	//response構造体に用意したデータを格納
	res := &model.User{
		Age: 20,
		CreatedAt: "2022-1-10",
		Faculty: 1,
		Gender: 1,
		Id: 1,
		Mail: "is0000@aaa.com",
		Name: "NN",
		UpdatedAt: "2022-2-10",
	}
	
	return http.StatusOK, res, nil
}

func (a *User) DeleteUser(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	
	return http.StatusOK, nil, nil
}
