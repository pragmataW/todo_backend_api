package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pragmataW/to-do/db"
	"github.com/pragmataW/to-do/instances"
)

func HandleLogin(c *fiber.Ctx) error {
	//* Connect to Db
	var	userInput	instances.LoginUser
	var	user		instances.User
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", db.DbUsername, db.DbPassword, db.DbHost, db.DbName))
	if err != nil {
		log.Println("Database Error: ", err)
		return c.Status(500).SendString(fmt.Sprintf("Database Error, database body cannot reachable %s", err))
	}
	defer db.Close()

	//* Get request body
	err = c.BodyParser(&userInput)
	if err != nil {
		log.Println("Json Error: ", err)
		return c.Status(400).SendString(fmt.Sprintf("Json Error, json body cannot reachable %s", err))
	}

	//* Compare with actual login data
	rows, err := db.Query("SELECT * FROM users WHERE username = ? AND password = ?", userInput.Username, userInput.Password)
	if err != nil{
		log.Println("Database Error: ", err)
		return c.Status(500).SendString(fmt.Sprintf("Database Error, database body cannot reachable %s", err))
	}
	if !rows.Next(){
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	rows.Scan(&user.Username, &user.Password, &user.Name, &user.Surname)

	//* Create token with custom claims
	claims := jwt.MapClaims{
		"Name": user.Name,
		"Surname": user.Surname,
		"Username": user.Username,
		"exp": time.Now().Add(time.Hour *72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenizedString, err := token.SignedString([]byte(instances.JwtPass))
	if err != nil{
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//* Set as cookie
	c.Cookie(&fiber.Cookie{
		Name: "Authentication",
		Value: tokenizedString,
		HTTPOnly: true,
		Expires: time.Now().Add(time.Hour * 72),
	})

	return c.SendStatus(http.StatusOK)
}