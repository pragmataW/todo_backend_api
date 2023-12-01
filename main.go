package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pragmataW/to-do/db"
	"github.com/pragmataW/to-do/routes"
)

func main(){
	app := fiber.New()

	routes.Listener(app)

	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil{
		fmt.Println("Error: ",  err)
	}
}

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Init()
	fmt.Println(db.DbUsername, db.DbPassword, db.DbHost, db.DbName)
	err = db.CreateUserTable()
	if err != nil{
		os.Exit(1)
	}
	err = db.CreateTodosTable()
	if err != nil{
		os.Exit(1)
	}
	
}