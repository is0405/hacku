package model

// Appeal defines model for Appeal.
type Appeal struct {
	Conditions       string `db:"conditions" json:"conditions"`
	Contents         string `db:"contents" json:"contents"`
	CreatedAt        string `db:"created_at" json:"createdAt"`
	Id               int    `db:"id" json:"id"`
	MaxParticipation int    `db:"max_participation" json:"maxParticipation"`
	Reward           string `db:"reward" json:"reward"`
	SubmitId         int    `db:"submit_id" json:"submitId"`
	Term             string `db:"term" json:"term"`
	Title            string `db:"title" json:"title"`
	UpdatedAt        string `db:"update_at" json:"updatedAt"`
}
