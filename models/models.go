package models

import (
	"time"

	"gorm.io/gorm"
)

type Lecture struct {
	gorm.Model
	name     string
	email    string
	password string
	class    []Class
}

type Class struct {
	gorm.Model
	lectureId   int
	name        string
	description string
	material    []Material
	student     []Student `gorm:"many2many:class_student"`
}

type Student struct {
	gorm.Model
	name     string
	email    string
	password string
}

type Assignment struct {
	gorm.Model
	description string
	due         time.Time
	score       int
}

type Material struct {
	gorm.Model
	name        string
	description string
	classId     int
}
