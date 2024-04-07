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

	DbUser = dbUser
	DbPass = dbPass
	DbName = dbName
	DbPort = dbPort
}
