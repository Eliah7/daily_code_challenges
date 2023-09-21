package main

import (
	"day3/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Day 3 Challenge")

	r := gin.Default()

	v1 := r.Group("/user")
	{
		v1.POST("/login", routes.Login)
		v1.POST("/signup", routes.SignUp)
		v1.GET("/users", routes.GetUsers)
	}
	// r.GET("/ping", routes.PingPost)
	// r.GET("/user/login", routes.Login)

	r.Run("localhost:8084")
}
