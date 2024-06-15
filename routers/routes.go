package routers

import (
	"github.com/gin-gonic/gin"
)






func SetUpRotes(router *gin.Engine){
	router.POST("/sign_up",SignUpRouter)
	router.POST("/login",LoginRouter)
}