package models

import (
	"time"
)

type User struct {
	Id           int       `json:"ID" gorm:"primary_key"`
	Name         string    `json:"name,omitempty"`
	Passwd       string    `json:"passwd,omitempty"`
	Prompt       string    `json:"Prompt,omitempty"`
	Answer       string    `json:"answer,omitempty"`
	Truename     string    `json:"truename,omitempty"`
	Idnumber     string    `json:"idnumber,omitempty"`
	Email        string    `json:"email,omitempty"`
	Mobilenumber string    `json:"mobilenumber,omitempty"`
	Province     string    `json:"province,omitempty"`
	City         string    `json:"city,omitempty"`
	Phonenumber  string    `json:"phonenumber,omitempty"`
	Address      string    `json:"address,omitempty"`
	Postalcode   string    `json:"postalcode,omitempty"`
	Gender       int       `json:"gender,omitempty"`
	Birthday     time.Time `json:"birthday,omitempty"`
	Creatime     time.Time `json:"creatime,omitempty"`
	Qq           string    `json:"qq,omitempty"`
	Passwd2      string    `json:"passwd2,omitempty"`
}
