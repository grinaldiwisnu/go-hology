package repository

import (
	"golang.org/x/crypto/bcrypt"
	"hology/database"
	"hology/models"
	"hology/models/user"
	"time"
)

func FetchAllUser() (result models.Response, code int, error error) {
	db := database.CreateConn()

	users := []user.User{}
	db.Find(&users)

	return models.Response{
		Status:  true,
		Message: "Success fetch users",
		Data:    users,
	}, 200, nil
}

func GetUserDetail(email, password string) (result models.Response, code int, error error) {
	db := database.CreateConn()

	u := user.User{}
	db.Where(&user.User{UserEmail: email}).Find(&u)

	hp, _ := Hash(password)

	if !IsSame(u.UserPassword, hp) {
		return models.Response{
			Status:  false,
			Message: "Password is wrong!",
			Data:    hp,
		}, 400, nil
	}

	return models.Response{
		Status:  true,
		Message: "Success get user",
		Data:    u,
	}, 200, nil
}

func CreateUser(usr user.User) (result models.Response, code int, error error) {
	db := database.CreateConn()

	hash, _ := Hash(usr.UserPassword)
	usr.UserPassword = hash
	usr.CreatedAt = time.Now()
	usr.UpdatedAt = time.Now()

	db.Create(usr)

	return models.Response{
		Status:  true,
		Message: "Success create user",
		Data:    usr,
	}, 200, nil
}

func Hash(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hashed), err
}

func IsSame(str string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}
