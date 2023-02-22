package Route

import (
	"fmt"
	"main/Init"
	"main/Model"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)

	hashedpass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error hashing password",
		})
		return
	}
	var user Model.User
	user = Model.User{Email: body.Email, Password: string(hashedpass)}
	result := Init.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Error creating user",
		})
		return
	}
	c.JSON(200, gin.H{})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	c.Bind(&body)

	var user Model.User
	Init.DB.Where("email = ?", body.Email).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error logging in",
		})
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	fmt.Println(tokenString)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Cannot Sign Token",
		})

	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)
}

func Validate(c *gin.Context) {
	users, _ := c.Get("user")
	c.JSON(200, gin.H{
		"logged in user": users,
	})
}
