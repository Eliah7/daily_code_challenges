package main

import (
	"day4/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Day4 Challenge")

	r := gin.Default()

	v1 := r.Group("/user")
	{
		v1.POST("/login", routes.Login)
		v1.POST("/singup", routes.SignUp)
		v1.GET("/user", routes.GetUsers)
	}

	r.Run("localhost:8084")
}
