package databaseutils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBClient struct {
	Username string
	Password string
	DBName string
	Opts string
}

var Conn *gorm.DB

func ConnectToDB(client DBClient){
	var err error

	Conn, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@/%s?%s",client.Username,client.Password,client.DBName,client.Opts))
	if err!=nil{
		panic(err.Error())
	}
	defer Conn.Close()
}