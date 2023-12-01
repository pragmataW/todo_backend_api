package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DbUsername	string
	DbPassword	string
	DbHost		string 
	DbName		string
)

func CreateUserTable() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", DbUsername, DbPassword, DbHost, DbName))
	if err != nil{
		log.Println(err)
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		username VARCHAR(20) PRIMARY KEY NOT NULL,
		password TEXT NOT NULL,
		name TEXT NOT NULL,
		surname TEXT NOT NULL
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CreateTodosTable() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", DbUsername, DbPassword, DbHost, DbName))
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}

	query := `
		CREATE TABLE IF NOT EXISTS todos (
			id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
			username VARCHAR(20) NOT NULL,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			FOREIGN KEY (username) REFERENCES users(username)
		);
	`

	_, err = db.Exec(query)
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}