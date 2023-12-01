package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/pragmataW/to-do/db"
	"github.com/pragmataW/to-do/instances"
)

func AddTodos(c *fiber.Ctx) error {
	var todo instances.Todos

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", db.DbUsername, db.DbPassword, db.DbHost, db.DbName))
	if err != nil{
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	defer db.Close()

	err = c.BodyParser(&todo)
	if err != nil{
		log.Println(err)
		return c.SendStatus(http.StatusBadRequest)
	}

	_, err = db.Exec("INSERT INTO todos (username, title, content) VALUES (?, ?, ?)", c.Locals("username"), todo.Title, todo.Content)
	if err != nil{
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func GetTodos(c *fiber.Ctx) error {
	username := c.Locals("username")
	var todos []instances.Todos

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", db.DbUsername, db.DbPassword, db.DbHost, db.DbName))
	if err != nil{
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, content FROM todos WHERE username = ?", username)
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}

	for rows.Next(){
		var todo instances.Todos
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Content)
		if err != nil {
			log.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	username := c.Locals("username")
	id := c.Params("id")
	var todo instances.Todos

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", db.DbUsername, db.DbPassword, db.DbHost, db.DbName))
	if err != nil{
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	defer db.Close()

	rows, err := db.Query("SELECT title, content FROM todos WHERE username = ? AND id = ?", username, id)
	if err !=  nil {
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	if !rows.Next(){
		return c.SendStatus(http.StatusBadRequest)
	}

	todo.Id, _ = strconv.Atoi(id)
	rows.Scan(&todo.Title, &todo.Content)
	return c.JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	todoId := c.Params("id")
	newTitle := c.Query("title")
	newContent := c.Query("content")
	username :=  c.Locals("username")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", db.DbUsername, db.DbPassword, db.DbHost, db.DbName))
	if err != nil{
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id FROM todos WHERE username = ? AND id = ?", username, todoId)
	if !rows.Next(){
		return c.SendStatus(http.StatusBadRequest)
	}

	if newTitle != ""{
		_, err = db.Exec("UPDATE todos SET title = ? WHERE username = ? AND id = ?", newTitle, username, todoId)
		if err != nil{
			log.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}
	}
	if newContent != ""{
		_, err = db.Exec("UPDATE todos SET content = ? WHERE username = ? AND id = ?", newContent, username, todoId)
		if err != nil{
			log.Println(err)
			return c.SendStatus(http.StatusInternalServerError)
		}
	}
	return c.SendStatus(http.StatusOK)
}

func DeleteTodo(c *fiber.Ctx) error {
	username := c.Locals("username")
	todoId := c.Params("id")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", db.DbUsername, db.DbPassword, db.DbHost, db.DbName))
	if err != nil{
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id FROM todos WHERE username = ? AND id = ?", username, todoId)
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	if !rows.Next(){
		return c.SendStatus(http.StatusBadRequest)
	}

	_, err = db.Exec("DELETE FROM todos WHERE username = ? AND id = ?", username, todoId)
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}