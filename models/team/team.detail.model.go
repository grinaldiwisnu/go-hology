package team

type TeamDetail struct {
	DetailTeamId           int    `json:"detail_team_id"`
	UserId                 int    `json:"user_id"`
	TeamId                 int    `json:"team_id"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
	DetailTeamIndentityPic string `json:"detail_team_indentity_pic"`
	DetailTeamProof        string `json:"detail_team_proof"`
}
