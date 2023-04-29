package controller

import (
	"MiniProject/config"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	config.InitDB()
	config.InitMigrate()
	return c.JSON(200, "OK")
}

func UpdateAllUsers(c echo.Context) error {
	return c.JSON(200, "OK")
}

func DeleteAllUsers(c echo.Context) error {
	return c.JSON(200, "OK")
}
