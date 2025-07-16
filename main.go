package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/mail"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `json:"user_name"`
	Email    string `json:"user_email"`
	Password string `json:"user_password"`
}

func CreatUser(ctx echo.Context) error {
	var user User

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "invalid body",
		})
	}

	err := validateUser(user)
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	hashPasword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {

		return ctx.JSON(500, echo.Map{
			"error": "error while hashing a password",
		})
	}

	user.Password = string(hashPasword)

	if err := saveTofile(user); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "user created successfully",
	})

}

func GetUser(ctx echo.Context) error {
	filename := ctx.Param("filename")
	user, err := ReadFromfile(filename)

	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": user,
	})
}

func validateUser(u User) error {
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return errors.New("all fields are required")

	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.New("invalid email format")
	}

	return nil
}

func saveTofile(u User) error {
	data, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("user-json", data, 0644)
}

func ReadFromfile(filename string) (User, error) {

	var user User

	bytes, err := os.ReadFile(filename)

	if err != nil {
		return user, err
	}

	err = json.Unmarshal(bytes, &user)

	return user, err
}

func main() {

	e := echo.New()

	e.POST("/creat", CreatUser)

	e.GET("/:filename", GetUser)
	fmt.Println("Server running on http://localhost:8080")

	e.Logger.Fatal(e.Start(":8080"))

}
