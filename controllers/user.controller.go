package controllers

import (
	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
	"hology/models"
	"hology/repository"
	"net/http"
)

func GetAll(context echo.Context) error {
	result, err := repository.FetchAllUser()

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
		"email":    []string{"required", "min:4", "max:20", "email"},
		"password": []string{"required", "min:8"},
	}

	val := govalidator.New(govalidator.Options{
		Request:         context.Request(),
		Rules:           rules,
		RequiredDefault: true,
	})

	e := val.Validate()

	if e != nil {
		return context.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: e.Encode(),
			Data:    nil,
		})
	}

	result, err := repository.GetUserDetail(context.FormValue("email"), context.FormValue("password"))

	if err != nil {
		return context.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return context.JSON(http.StatusOK, result)
}
