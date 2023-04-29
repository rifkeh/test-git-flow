package config

import (
	"MiniProject/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "rifkhi",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "LMS",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"], config["DB_Name"])
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.Lecture{})
	DB.AutoMigrate(&models.Class{})
	DB.AutoMigrate(&models.Student{})
	DB.AutoMigrate(&models.Material{})
	DB.AutoMigrate(&models.Assignment{})
}
