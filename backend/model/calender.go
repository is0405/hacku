package model

type Calender struct {
	Id     int    `db:"id" json:"carender_id"`
	Date   string `db:"date" json:"date"`
	Time   string `db:"time" json:"time"`
	BookOk bool   `db:"book_ok" json "book_ok"`
	Iam    bool   `db:"iam" json "iam"` 
}
