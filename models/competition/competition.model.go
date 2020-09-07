package competition

type Competition struct {
	CompetitionId          int    `json:"competition_id"`
	CompetitionName        string `json:"competition_name"`
	CompetitionDescription string `json:"competition_description"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}
