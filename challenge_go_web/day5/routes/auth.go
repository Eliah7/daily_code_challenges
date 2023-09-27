package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	var body struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, GetHttpStatus(http.StatusBadRequest))
		return
	}

	var user User
	var userFound bool
	for _, user_ := range users {
		if user_.UserName == body.UserName {
			user = user_
			userFound = true
		}
	}

	if !userFound {
		c.JSON(http.StatusNotFound, GetHttpStatus(http.StatusNotFound))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(body.Password))
	fmt.Println(err.Error())
	fmt.Println(user.password, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, GetHttpStatus(http.StatusUnauthorized))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"id":  user.ID, // exposing internals
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString(os.Getenv("SECRET"))

	if err != nil {
		c.JSON(http.StatusBadGateway, GetHttpStatus(http.StatusNotFound))
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "localhost", false, true)
}

func SignUp(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return // return an error for error to signup
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
