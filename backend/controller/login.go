package controller

import (
	// "encoding/json"
	// "errors"
	"net/http"

	// "github.com/is0405/hacku/httputil"
	// "github.com/is0405/hacku/model"
	// "github.com/is0405/hacku/repository"
	// "github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

type Login struct {
	db           *sqlx.DB
	jwtSecretKey []byte
}

func NewLogin(db *sqlx.DB, jwtSecretKey []byte) *Login {
	return &Login{db: db, jwtSecretKey: jwtSecretKey}
}

// PostLoginJSONBody defines parameters for PostLogin.
type PostLoginJSONBody struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	UserId    int    `json:"user_id"`
}

func (a *Login) Login(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	//response構造体に用意したデータを格納
	res := LoginResponse{
		Token: "",
		UserId: 0,
	}

	return http.StatusOK, res, nil
}
