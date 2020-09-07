package repository

import (
	"hology/database"
	"hology/models"
	"hology/models/user"
)

func FetchAllUser() (models.Response, error) {
	var obj user.User
	var arrobj []user.User
	var res models.Response

	con := database.CreateConn()

	results, err := con.Query("SELECT * FROM users ORDER BY user_id")
	defer results.Close()

	if err != nil {
		return res, err
	}

	for results.Next() {
		err = results.Scan(&obj.UserId, &obj.InstitutionId, &obj.UserFullname, &obj.UserEmail, &obj.UserName, &obj.UserPassword, &obj.UserGender, &obj.UserBirthdate, &obj.CreatedAt, &obj.UpdatedAt, &obj.TeamId)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	return models.Response{
		Status:  true,
		Message: "Success fetch users",
		Data:    arrobj,
	}, nil
}

func GetUserDetail(email, password string) (models.Response, error) {
	var obj user.User
	var arrobj []user.User
	var res models.Response

	con := database.CreateConn()

	results, err := con.Query("SELECT * FROM users WHERE user_email = " + email)
	defer results.Close()

	if err != nil {
		return res, err
	}

	for results.Next() {
		err = results.Scan(&obj.UserId, &obj.InstitutionId, &obj.UserFullname, &obj.UserEmail, &obj.UserName, &obj.UserPassword, &obj.UserGender, &obj.UserBirthdate, &obj.CreatedAt, &obj.UpdatedAt, &obj.TeamId)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	return models.Response{
		Status:  true,
		Message: "Success auth users",
		Data:    arrobj,
	}, nil
}
