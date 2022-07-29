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

type Hired struct {
	db           *sqlx.DB
}

func NewHired(db *sqlx.DB) *Hired {
	return &Hired{db: db}
}


func (a *Hired) PostHired(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}

func (a *Hired) DeleteHired(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}

