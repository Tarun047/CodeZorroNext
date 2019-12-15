package repository

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"models"
)

type CodeRepository struct {
	Conn *gorm.DB
}

func (repo *CodeRepository) GetById(id uuid.UUID) models.Code {
	var retVal models.Code
	repo.Conn.Where("id = ?",id).First(&retVal)
	return retVal
}

func (repo *CodeRepository) Save(code *models.Code){
	repo.Conn.Create(&code)
}