package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
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
func (contact *Code) Validate() (map[string] interface{}, bool) {

	if contact.Source == "" {
		return u.Message(false, "Code name should be on the payload"), false
	}

	if contact.Language == "" {
		return u.Message(false, "Language number should be on the payload"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (contact *Code) Create() (map[string] interface{}) {

	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) (*Code) {

	contact := &Code{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) ([]*Code) {

	contacts := make([]*Code, 0)
	err := GetDB().Table("codes").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}