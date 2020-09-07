package team

type Team struct {
	TeamId           int    `json:"team_id"`
	InstitutionId    int    `json:"institution_id"`
	CompetitionId    int    `json:"competition_id"`
	TeamName         string `json:"team_name"`
	TeamPaymentProof string `json:"team_payment_proof"`
	TeamLead         int    `json:"team_lead"`
	TeamStatus       int    `json:"team_status"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	TeamJoinUrl      string `json:"team_join_url"`
}
