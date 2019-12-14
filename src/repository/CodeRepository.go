package repository

import (
	"github.com/jinzhu/gorm"
	"models"
)

type CodeRepository struct {
	Conn *gorm.DB
}

func (repo *CodeRepository) GetById(id int) models.Code {
	var retVal models.Code
	repo.Conn.First(&retVal,id)
	return retVal
}

func (repo *CodeRepository) Save(code models.Code){
	repo.Conn.Create(&code)
}