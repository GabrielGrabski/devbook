package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Dns  = ""
	Port = 0
)

func GetEnvVars() {
	var err error
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		Port = 5000
	}

	dbUser, dbPassword, dbName := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")

	Dns = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&Loc=Local", dbUser, dbPassword, dbName)
}
