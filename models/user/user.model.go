package user

import "time"

type User struct {
	UserId        int       `json:"user_id"`
	InstitutionId int       `json:"institution_id"`
	UserFullname  string    `json:"user_fullname"`
	UserEmail     string    `json:"user_email"`
	UserName      string    `json:"user_name"`
	UserPassword  string    `json:"-"`
	UserGender    string    `json:"user_gender"`
	UserBirthdate string    `json:"user_birthdate"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	TeamId        string    `json:"team_id"`
}
