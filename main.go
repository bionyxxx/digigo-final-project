package main

import (
	"Final_Project/configs"
	"Final_Project/routes"
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
	var PORT = ":3000"
	err := routes.ApiInit().Run(PORT)
	if err != nil {
		panic(err)
	}
}
