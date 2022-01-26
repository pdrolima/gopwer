package models

import (
	"github.com/webmasterdro/gopwer/database"
	"time"
)

type AddCash struct {
}

type User struct {
	ID           int       `json:"ID" gorm:"primary_key"`
	Name         string    `json:"name,omitempty;" validate:"required,min=3,max=32"`
	Passwd       string    `json:"passwd,omitempty;",validate:"required,min=6,max=32"`
	Prompt       string    `json:"Prompt,omitempty;"validate:"max=32"`
	Answer       string    `json:"answer,omitempty;"validate:"max=32"`
	Truename     string    `json:"truename,omitempty;"validate:"max=32"`
	Idnumber     string    `json:"idnumber,omitempty;"validate:"max=32"`
	Email        string    `json:"email,omitempty;"validate:"required,max=64"`
	Mobilenumber string    `json:"mobilenumber,omitempty"`
	Province     string    `json:"province,omitempty"`
	City         string    `json:"city,omitempty"`
	Phonenumber  string    `json:"phonenumber,omitempty"`
	Address      string    `json:"address,omitempty"`
	Postalcode   string    `json:"postalcode,omitempty"`
	Gender       int       `json:"gender,omitempty;" validate:"min=0,max=1"`
	Birthday     time.Time `json:"birthday,omitempty"`
	Creatime     time.Time `json:"creatime,omitempty"`
	Qq           string    `json:"qq,omitempty"`
	Passwd2      string    `json:"passwd2,omitempty"`
}

func GetUserByUserName(username string) (*User, error) {
	user := User{}

	err := database.DB.Where("name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	err := database.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func AddUserCash(userid int, cash int) error {
	err := database.DB.Exec("call usecash(?, ?, ? ,?, ?, ?, ?, @error)", userid, 1, 0, 1, 0, cash*100, 1).Error
	if err != nil {
		return err
	}
	return nil
}
