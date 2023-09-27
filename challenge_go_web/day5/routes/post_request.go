package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	fmt.Println(GetHttpStatus(http.StatusOK))
	c.JSON(http.StatusOK, GetHttpStatus(http.StatusOK))
}

func GetPosts(c *gin.Context) {
	fmt.Println(GetHttpStatus(http.StatusOK))
	c.JSON(http.StatusOK, GetHttpStatus(http.StatusOK))
}
