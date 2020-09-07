package routes

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"hology/controllers"
	"hology/models"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, echo.Version)
	})

	e.GET("/users", controllers.GetAll)
	e.POST("/auth", controllers.Login)
	e.GET("/pwd", func(context echo.Context) error {
		pwd, err := bcrypt.GenerateFromPassword([]byte("grinaldi"), bcrypt.DefaultCost)

		if err != nil {

		}

		return context.JSON(http.StatusOK, models.Response{
			Status:  true,
			Message: "generated",
			Data:    pwd,
		})
	})

	return e
}
