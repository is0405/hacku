package model

type Recruitment struct {
	Conditions       string `db:"conditions" json:"conditions"`
	Contents         string `db:"contents" json:"contents"`
	CreatedAt        string `db:"created_at" json:"createdAt"`
	Id               int    `db:"id" json:"id"`
	MaxParticipation int    `db:"max_participation" json:"maxParticipation"`
	Reward           string `db:"reward" json:"reward"`
	SubmitId         int    `db:"submit_id" json:"submitId"`
	StartRecruitmentPeriod     string `db:"start_recruitment_period" json:"startRecruitmentPeriod"`
	FinishRecruitmentPeriod    string `db:"finish_recruitment_period" json:"finishRecruitmentPeriod"`
	StartImplementationPeriod  string `db:"start_implementation_period" json:"startImplementationPeriod"`
	FinishImplementationPeriod string `db:"finish_implementation_period" json:"finishImplementationPeriod"`
	Gender           int `db:"sex" json: sex"`
	Title            string `db:"title" json:"title"`
	MinAge           int `db:"min_age" json: minAge"`
	MaxAge           int `db:"max_age" json: maxAge"`
	UpdatedAt        string `db:"update_at" json:"updatedAt"`
}
