package models

import "github.com/jinzhu/gorm"

type Interview struct {
	gorm.Model
	interviewerId uint
	intervieweeId uint
	codeChannel chan string
}







