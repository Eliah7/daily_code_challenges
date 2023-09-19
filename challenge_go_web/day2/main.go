package main

import (
	"day2/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Day 2 challenge")

	r := gin.Default()
	v1 := r.Group("/user")
	{
		v1.GET("/login", routes.Login)
		v1.GET("/signUp", routes.SignUp)
	}
	// r.GET("/ping", routes.PingPost)
	// r.GET("/user/login", routes.Login)

	r.Run("localhost:8084")

}
