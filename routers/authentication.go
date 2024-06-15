package routers

import (

	"github.com/aglili/go-expense/models"
	"github.com/aglili/go-expense/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)






func SignUpRouter(c *gin.Context){
	var body struct {
        Email    string `json:"email" binding:"required"`
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }


	if c.Bind(&body) != nil {
		c.JSON(400,gin.H{"error":"Invalid request"})
		return
	}


	hash,err := bcrypt.GenerateFromPassword([]byte(body.Password),bcrypt.DefaultCost)

	if err != nil{
		c.JSON(500,gin.H{"error":"Internal server error"})
		return
	}

	user := models.User{
		Email: body.Email,
		Username: body.Username,
		Password: string(hash),
	}


	err = services.CreateUser(user)
	if err != nil{
		c.JSON(400,gin.H{"error":"Failed to create user"})
		return
	}

	c.JSON(200,gin.H{"message":"User created successfully"})

}



func LoginRouter(c *gin.Context){
	var body struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}


	if c.Bind(&body) != nil {
		c.JSON(400,gin.H{"error":"Invalid request"})
		return
	}

	token,err := services.LoginUser(body.Email,body.Username,body.Password)
	if err != nil{
		c.JSON(400,gin.H{"error":"Invalid credentials"})
		return
	}

	c.JSON(200,gin.H{"token":token})
	

}