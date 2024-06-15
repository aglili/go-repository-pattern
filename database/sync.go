package database

import "github.com/aglili/go-expense/models"


func SyncDatabase (){
	DB.AutoMigrate(&models.User{})
}