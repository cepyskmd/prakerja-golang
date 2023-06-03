package routes

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/tasks", controllers.GetTaskController)
	e.POST("/tasks", controllers.InsertTaskController)
	e.PUT("/tasks", controllers.UpdateTaskController)
	e.DELETE("/tasks/:id", controllers.DeleteTaskController)
	return e
}
