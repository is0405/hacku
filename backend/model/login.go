package model

type LoginBody struct {
	Mail     string `db:"mail" json:"mail"`
	Password string `db:"password" json:"password"`
}
