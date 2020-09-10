package routes

import (
	"github.com/labstack/echo"
	"hology/controllers"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, echo.Version)
	})

	e.GET("/users", controllers.GetAll)
	e.POST("/auth", controllers.Login)
	e.POST("/register", controllers.Register)

	return e
}
