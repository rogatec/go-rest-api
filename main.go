package main

import (
	"os"

	_ "github.com/joho/godotenv"
)

func main() {
	LoadEnv()

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
