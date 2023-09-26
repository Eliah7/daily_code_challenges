package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"id"`
	UserName     string `json:"user_name"`
	user_type_id string `json:"user_type_id"`
	password     string `json:"password"`
}

var users = []User{
	{ID: "1", UserName: "Blue Train", password: "John Coltrane"},
	{ID: "2", UserName: "Jeru", password: "Gerry Mulligan"},
	{ID: "3", UserName: "Sarah Vaughan and Clifford Brown", password: "Sarah Vaughan"},
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Login(c *gin.Context) {
	fmt.Println(GetHttpStatus(http.StatusOK))
	c.JSON(http.StatusOK, GetHttpStatus(http.StatusOK))
}

func SignUp(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.password), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to generate password")
	}

	newUser.password = string(hashedPassword)
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
