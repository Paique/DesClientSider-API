package data

import (
	"log"
	"os"
)

var (
	AppPort string

	DbUser string
	DbPass string
	DbName string
	DbPort string
	DbHost string
)

func init() {
	appPort, isAppPortSet := os.LookupEnv("DCS_LISTEN_PORT")

	if !isAppPortSet {
		log.Println("DCS_LISTEN_PORT not set, defaulting to 8080")
		appPort = "8080"
	}

	AppPort = appPort

	dbUser, isDbUserSet := os.LookupEnv("DB_USERNAME")
	dbPass, isDbPassSet := os.LookupEnv("DB_PASSWORD")
	dbName, isDbNameSet := os.LookupEnv("DB_NAME")
	dbPort, isDbPortSet := os.LookupEnv("DB_PORT")
	dbHost, isDbHostSet := os.LookupEnv("DB_HOST")

	if !isDbPassSet {
		log.Panicln("DB_PASSWORD environment variable not set")
	}

	if !isDbUserSet {
		log.Panicln("DB_USERNAME environment variable not set")
	}

	if !isDbNameSet {
		log.Println("DB_NAME environment variable not set, defaulting to dcs")
		dbName = "dcs"
	}

	if !isDbPortSet {
		log.Println("DB_PORT environment variable not set, defaulting to 3306")

		dbPort = "3306"
	}

	if !isDbHostSet {
		log.Println("DB_HOST environment variable not set")
		dbHost = "127.0.0.1"
	}

	DbUser = dbUser
	DbPass = dbPass
	DbName = dbName
	DbPort = dbPort
	DbHost = dbHost
}
