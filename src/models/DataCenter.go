package models


import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

	//Conn, err = gorm.Open("mysql",
	//	fmt.Sprintf("%s:%s@/%s?%s",client.Username,client.Password,client.DBName,client.Opts))
	Conn,err = gorm.Open("sqlite3","test.db")
	//Turn off debug in production
	Conn.LogMode(true)
	if err!=nil{
		panic(err.Error())
	}
}

func MakeMigrations()  {
	Conn.AutoMigrate(&Code{},&Token{},&Account{})

}

func GetDB() *gorm.DB {
	return Conn
}



