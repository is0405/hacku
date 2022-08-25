package model

type Recruitment struct {
	Conditions       string `db:"conditions" json:"conditions"`
	Contents         string `db:"contents" json:"contents"`
	CreatedAt        string `db:"created_at" json:"createdAt"`
	Id               int    `db:"id" json:"id"`
	MaxParticipation int    `db:"max_participation" json:"maxParticipation"`
	Reward           string `db:"reward" json:"reward"`
	SubmitId         int    `db:"submit_id" json:"submitId"`
	Period           string `db:"period" json:"period"`
	Gender           int `db:"gender" json: sex"`
	Title            string `db:"title" json:"title"`
	MinAge           int `db:"min_age" json: minAge"`
	MaxAge           int `db:"max_age" json: maxAge"`
	UpdatedAt        string `db:"updated_at" json:"updatedAt"`
}
