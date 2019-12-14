package main

import (
	"databaseutils"
	"github.com/jinzhu/gorm"
	"models"
	"repository"
)


func setup() {
	client := databaseutils.DBClient{
		Username: "root",
		Password: "tarun123",
		DBName:   "codezorro",
	}
	databaseutils.ConnectToDB(client)
}

func main() {
	setup()
	test:=repository.CodeRepository{Conn:databaseutils.Conn}
	obj: = models.Code{
		Source:   "Hello World",
		Language: "Java",
	}
	
}
