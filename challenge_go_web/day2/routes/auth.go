package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println(GetHttpStatus(http.StatusOK))
	c.JSON(http.StatusOK, GetHttpStatus(http.StatusOK))
}

func SignUp(c *gin.Context) {
	fmt.Println(GetHttpStatus(http.StatusOK))
	c.JSON(http.StatusOK, GetHttpStatus(http.StatusOK))
}
