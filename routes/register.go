package routes

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/pragmataW/to-do/db"
	"github.com/pragmataW/to-do/instances"
)

func usernameControl(user instances.User, db *sql.DB) error {
	length := len(user.Username)
	if length > 20 || length < 8{
		unameErr := instances.UsernameError{Message: "Username must be maximum 20 character or minimum 8 character",}
		return unameErr
	}

	row, err := db.Query("SELECT * FROM users WHERE username = ?", user.Username)
	if err != nil{
		return err
	}

	if row.Next(){
		unameErr := instances.UsernameError{Message: "Username must be unique",}
		return unameErr
	}

	return nil
}

func  passwordControl(user instances.User) error {
	length := len(user.Password)
	if length < 8{
		passErr := instances.PasswordError{Message: "Password must be minimum 8 character",}
		return passErr
	}
	return nil
}

func registerControl(user instances.User, db *sql.DB) error {
	err := usernameControl(user, db)
	if err != nil{
		return err
	}
	err = passwordControl(user)
	if err != nil{
		return err
	}
	return nil
}

func HandleRegister(c *fiber.Ctx) error {
	var user instances.User
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", db.DbUsername, db.DbPassword, db.DbHost, db.DbName))
	if err != nil{
		log.Println("Database Error: ", err)
		return c.Status(500).SendString(fmt.Sprintf("Database Error, database body cannot reachable %s", err))
	}
	defer db.Close()

	err = c.BodyParser(&user)
	if err != nil{
		log.Println("Json Error: ", err)
		return c.Status(400).SendString(fmt.Sprintf("Json Error, json body cannot reachable %s", err))
	}

	err = registerControl(user, db)
	switch err.(type){
	case instances.UsernameError, instances.PasswordError:
		log.Printf("%s", err)
		return c.Status(500).SendString(fmt.Sprintf("%s", err))
	case nil:
		//contiune
	default:
		log.Printf("%s", err)
		return c.Status(500).SendString(fmt.Sprintf("%s", err))
	}

	_, err = db.Exec("INSERT INTO users (username, password, name, surname) VALUES (?, ?, ?, ?)", user.Username, user.Password, user.Name, user.Surname)
	if err != nil{
		log.Println("Database Error", err)
		return c.Status(500).SendString(fmt.Sprintf("Database Error: %s", err))
	}

	log.Printf("%s has been registered", user.Username)
	return c.Status(201).SendString(fmt.Sprintf("%s has been registered", user.Username))
}