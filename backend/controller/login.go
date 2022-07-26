package controller

import (
	"encoding/json"
	// "errors"
	"net/http"
	"time"
	"fmt"

	// "github.com/is0405/hacku/httputil"
	"github.com/is0405/hacku/model"
	"github.com/is0405/hacku/repository"
	"github.com/is0405/hacku/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

type Login struct {
	db           *sqlx.DB
	jwtSecretKey []byte
}

func NewLogin(db *sqlx.DB, jwtSecretKey []byte) *Login {
	return &Login{db: db, jwtSecretKey: jwtSecretKey}
}

type LoginResponse struct {
	Token     string `json:"token"`
}

func (a *Login) Login(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	lb := &model.LoginBody{}

	err := json.NewDecoder(r.Body).Decode(&lb);
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	
	if !util.MailCheck(lb.Mail) {
		return http.StatusBadRequest, nil, err
	}

	lb.Password = util.HashGenerateSha256(lb.Password)
	fmt.Println(lb.Mail)
	fmt.Println(lb.Password)
	userId, err := repository.CheckLogin(a.db, lb)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	res, err := TokenGenerate(userId, a.jwtSecretKey)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

func TokenGenerate(userId int, key []byte) (LoginResponse, error) {
	// jwtの作成
	claims := model.Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //トークン発行から24時間で使えなくなる.
			Issuer:    "com.hacku",
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)

	if err != nil {
		return LoginResponse{}, err
	}
	
	//response構造体に用意したデータを格納
	res := LoginResponse{
		Token: signedToken,
	}
	return res, nil
}
