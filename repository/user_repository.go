package repository

import (
	"github.com/aglili/go-expense/database"
	"github.com/aglili/go-expense/models"
)


type UserRepostory interface{
	CreateUser(user models.User) error
	GetExistingUser(email,username string) (models.User, error)
}




func CreateUser(user models.User) error{
	err := database.DB.Create(&user).Error
	if err != nil{
		return err
	}
	return nil
}




func GetExistingUser(email,username string) (models.User, error){
	var user models.User
	err := database.DB.Where("email = ? OR username = ?", email, username).First(&user).Error
	if err != nil{
		return user, err
	}
	return user, nil
}
	