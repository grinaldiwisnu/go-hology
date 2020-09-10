package controllers

import (
	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
	"hology/models"
	"hology/models/user"
	"hology/repository"
	"net/http"
	"strconv"
)

func GetAll(context echo.Context) error {
	result, _, err := repository.FetchAllUser()

	if err != nil {
		return context.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return context.JSON(http.StatusOK, result)
}

func Login(context echo.Context) error {
	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required", "min:8"},
	}

	v := govalidator.New(govalidator.Options{
		Request:         context.Request(),
		Rules:           rules,
		RequiredDefault: true,
	})

	e := v.Validate()

	if e.Encode() != "" {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "Input error",
			Data:    nil,
		})
	}

	result, code, err := repository.GetUserDetail(context.FormValue("email"), context.FormValue("password"))

	if err != nil {
		return context.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	} else if code != 200 {
		return context.JSON(code, result)
	}

	return context.JSON(http.StatusOK, result)
}

func Register(context echo.Context) error {
	rules := govalidator.MapData{
		"email":       []string{"required", "email"},
		"password":    []string{"required", "min:8"},
		"fullname":    []string{"required"},
		"gender":      []string{"required"},
		"birthdate":   []string{"required"},
		"institution": []string{"required"},
	}

	v := govalidator.New(govalidator.Options{
		Request:         context.Request(),
		Rules:           rules,
		RequiredDefault: true,
	})

	e := v.Validate()

	if e.Encode() != "" {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "Input error",
			Data:    nil,
		})
	}

	inst, _ := strconv.Atoi(context.FormValue("institution"))

	usr := user.User{
		InstitutionId: inst,
		UserFullname:  context.FormValue("fullname"),
		UserEmail:     context.FormValue("email"),
		UserName:      context.FormValue("fullname"),
		UserPassword:  context.FormValue("password"),
		UserGender:    context.FormValue("gender"),
		UserBirthdate: context.FormValue("birthdate"),
		TeamId:        "",
	}

	res, code, err := repository.CreateUser(usr)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	} else if code != 200 {
		return context.JSON(code, res)
	}

	return context.JSON(http.StatusOK, res)
}
