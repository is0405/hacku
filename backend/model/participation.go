package model

type Participation struct 
{
	Id        int `db:"id" json:"id"`
	AppealId  int `db:"appeal_id" json:"appealId"`
	UserId    int `db:"user_id" json:"userId"`
}
