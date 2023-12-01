package endpoints

import (
	"fmt"
	"os"
)

var JwtPass = os.Getenv("SECRET")

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

type Todos struct {
	Id			int		`json:"id"`
	Title		string	`json:"title"`
	Content		string	`json:"content"`
}

type UsernameError struct {
	Message string
}

type PasswordError struct {
	Message string
}

func (e UsernameError) Error() string {
	return fmt.Sprintf("Username Error: %s", e.Message)
}

func (e PasswordError) Error() string {
	return fmt.Sprintf("Password Error: %s", e.Message)
}