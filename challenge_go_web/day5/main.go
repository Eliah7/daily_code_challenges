package main

import (
	"day5/inits"
	"day5/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
}

func main() {
	fmt.Println("Day5 Challenge !")
	r := gin.Default()

	v1 := r.Group("/user")
	{
		v1.POST("/login", routes.Login)
		v1.POST("/singup", routes.SignUp)
		v1.GET("/users", routes.GetUsers)
	}

	r.Run("localhost:8084")
}
