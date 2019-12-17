package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"os"
	"strconv"
	u "utils"
)

type Code struct {
gorm.Model
Source   string `json:"source"`
Language string `json:"language"`
UserId   uint   `json:"user_id"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/
func (code *Code) Validate() (map[string] interface{}, bool) {

	if code.Source == "" {
		return u.Message(false, "Code name should be on the payload"), false
	}

	if code.Language == "" {
		return u.Message(false, "Language number should be on the payload"), false
	}

	if code.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (code *Code) Create() map[string] interface{} {

	if resp, ok := code.Validate(); !ok {
		return resp
	}

	GetDB().Create(code)

	resp := u.Message(true, "success")
	resp["code"] = code
	return resp
}

func (code *Code) SaveToFile(nameChannel *chan string){
	var ext string
	switch code.Language {
	case "C": ext =".c"
	case "CPP": ext = ".cpp"
	case "PYTH": ext = ".py"
	case "JAVA": ext=".java"
	}
	os.Mkdir("codes",0755)
	f,err:=ioutil.TempFile("codes",strconv.Itoa(int(code.UserId))+"*"+ext)
	if err!=nil {
		panic(err.Error())
	}
	f.WriteString(code.Source)
	*nameChannel<-f.Name()
}

func (code *Code) Run() (string,bool){
	nameChannel := make(chan string)
	go code.SaveToFile(&nameChannel)
	fname := <-nameChannel
	return RunSpecificCode(code,fname)
}

func GetCode(id uint) *Code {

	code := &Code{}
	err := GetDB().Table("codes").Where("id = ?", id).First(code).Error
	if err != nil {
		return nil
	}
	return code
}

func GetCodes(user uint) []*Code {

	codes := make([]*Code, 0)
	err := GetDB().Table("codes").Where("user_id = ?", user).Find(&codes).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return codes
}