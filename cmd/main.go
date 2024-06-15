package main

import (
	"github.com/aglili/go-expense/config"
	"github.com/aglili/go-expense/database"
	"github.com/aglili/go-expense/routers"
	"github.com/gin-gonic/gin"
)





func init(){
	config.LoadVariables()
	database.ConnectToDatabase()
	database.SyncDatabase()
}





func main()  {
	

	router := gin.Default()


	routers.SetUpRotes(router)

	router.Run(":8000")





}