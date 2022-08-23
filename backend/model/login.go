package model

type LoginBody struct {
	Mail     string `db:"mail" json:"mailaddress"`
	Password string `db:"password" json:"password"`
}
