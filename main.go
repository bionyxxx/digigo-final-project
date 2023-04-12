package main

import (
	"Final_Project/configs"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = configs.InitDatabase()
	if err != nil {
		panic(err)
	}
}

func main() {
	StartApp()
}
