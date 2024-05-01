package routes

import (
	"net/http"
	"redtower/service/database"
	"redtower/service/database/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
	
	return err == nil
}

func Signup(c echo.Context) error {
	password := c.FormValue("password")	
	email := c.FormValue("email")
	username := c.FormValue("username")

	hashedPassword, err := HashPassword(password)

	if err != nil {
		c.Logger().Fatal("Password Hashing Failed")
	}

	user := models.User{Username: username, Email: email, PasswordHash: hashedPassword}

	database.DB.Create(&user)

	return c.String(http.StatusCreated, "User created")
}

func Login(c echo.Context) error {
	var user models.User	

	username := c.FormValue("username")
	password := c.FormValue("password")

	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.String(http.StatusBadRequest, "User Not Found")
	}

	result := CheckPasswordHash(password, user.PasswordHash)

	if !result {
		return c.String(http.StatusForbidden, "Password or Username is incorrect")
	}

	return c.JSON(http.StatusOK, user)
}