package models

import (
	"github.com/webmasterdro/gopwer/database"
	"github.com/webmasterdro/gopwer/entity"
)

func GetUserByUserName(username string) (*entity.User, error) {
	user := entity.User{}

	err := database.DB.Where("name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *entity.User) error {
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
