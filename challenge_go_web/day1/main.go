package main

import (
	"github.com/gin-gonic/gin"
	"day1/routes" 
)

func main(){
	r := gin.Default()

	r.GET("/ping", routes.PingPost)
	r.GET("/user/login", routes.Login)
	
	r.Run("localhost:8084")
}