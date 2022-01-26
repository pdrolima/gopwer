package models

import (
	"time"
)

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
	Gender       int       `json:"gender,omitempty"`
	Birthday     time.Time `json:"birthday,omitempty"`
	Creatime     time.Time `json:"creatime,omitempty"`
	Qq           string    `json:"qq,omitempty"`
	Passwd2      string    `json:"passwd2,omitempty"`
}
