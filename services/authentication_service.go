package services

import (
	"time"

	"github.com/aglili/go-expense/config"
	"github.com/aglili/go-expense/models"
	"github.com/aglili/go-expense/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)



func CreateUser(user models.User) error{
	err := repository.CreateUser(user)
	if err!= nil{
		return err
	}
	return nil
}



func LoginUser(email,username,password string) (string,error){
	user,err := repository.GetExistingUser(email,username)
	if err != nil{
		return "",err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil{
		return "",err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},

	)

	tokenString,err := token.SignedString([]byte(config.GetJWTSecret()))
	if err != nil{
		return "",err
	}

	return tokenString,nil

}

	



