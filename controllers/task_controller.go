package controllers

import (
	"myapp/configs"
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTaskController(c echo.Context) error {
	var tasks []models.Task
	var response models.BaseResponse

	result := configs.DB.Find(&tasks)

	if result.RowsAffected == 0 {
		response.Message = "Please Create a Task"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Message = "Success Create Data"
	response.Data = tasks
	return c.JSON(http.StatusOK, response)
}

func InsertTaskController(c echo.Context) error {
	var task models.Task
	var response models.BaseResponse

	c.Bind(&task)
	result := configs.DB.Create(&task)

	if result.Error != nil {
		response.Message = "Failed. Server Error"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Success Insert Data"
	response.Data = task
	return c.JSON(http.StatusOK, response)
}

func UpdateTaskController(c echo.Context) error {
	var task models.Task
	var response models.BaseResponse

	c.Bind(&task)
	result := configs.DB.Model(models.Task{}).Where("id = ?", task.Id).Updates(models.Task{Task: task.Task, IsDone: task.IsDone})

	if result.Error != nil {
		response.Message = "Failed. Server Error"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Success Update Data"
	response.Data = task
	return c.JSON(http.StatusOK, response)
}

func DeleteTaskController(c echo.Context) error {
	var response models.BaseResponse

	var id = c.Param("id")
	result := configs.DB.Delete(&models.Task{}, id)

	if result.Error != nil {
		response.Message = "Failed. Server Error"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Success Delete Data"
	return c.JSON(http.StatusOK, response)
}
