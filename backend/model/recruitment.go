package model

type Recruitment struct {
	Conditions       string `db:"conditions" json:"conditions"`
	Contents         string `db:"contents" json:"content"`
	CreatedAt        string `db:"created_at" json:"createdAt"`
	Id               int    `db:"id" json:"id"`
	MaxParticipation int    `db:"max_participation" json:"maxSubjects"`
	Reward           string `db:"reward" json:"reward"`
	SubmitId         int    `db:"submit_id" json:"submitId"`
	Period           string `db:"period" json:"period"`
	Gender           int    `db:"gender" json:"sex"`
	Title            string `db:"title" json:"title"`
	MinAge           int    `db:"min_age" json:"minAge"`
	MaxAge           int    `db:"max_age" json:"maxAge"`
	UpdatedAt        string `db:"updated_at" json:"updatedAt"`
}

type ReqRecruitment struct {
	Conditions       string `db:"conditions" json:"conditions"`
	Contents         string `db:"contents" json:"content"`
	CreatedAt        string `db:"created_at" json:"createdAt"`
	Id               int    `db:"id" json:"id"`
	MaxParticipation int    `db:"max_participation" json:"maxSubjects"`
	Reward           string `db:"reward" json:"reward"`
	SubmitId         int    `db:"submit_id" json:"submitId"`
	Period           string `db:"period" json:"period"`
	Gender           int    `db:"gender" json:"sex"`
	Title            string `db:"title" json:"title"`
	MinAge           int    `db:"min_age" json:"minAge"`
	MaxAge           int    `db:"max_age" json:"maxAge"`
	UpdatedAt        string `db:"updated_at" json:"updatedAt"`
	DateList         []Calender `json:"dateList"`
}
