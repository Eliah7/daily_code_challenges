package routes

import "github.com/gin-gonic/gin"
import "fmt"
import "net/http"

func PingPost(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Posts are available",
	})
}

func Login(c *gin.Context){
	fmt.Println(GetHttpStatus(http.StatusOK))
	c.JSON(http.StatusOK, gin.H{
		"status": GetHttpStatus(http.StatusOK),
	})
}