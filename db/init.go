package db

import "os"

func Init() {
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost =  os.Getenv("DB_HOST")
	DbName =  os.Getenv("DB_NAME")
}