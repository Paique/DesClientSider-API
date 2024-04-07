package data

import (
	"database/sql"
	"os"
)

var (
	AppPort string

	DbUser string
	DbPass string
	DbName string
	DbPort string
	DbHost string

	DbConn    *sql.DB
	ConnTries int = 0
)

func InitVariables() {

	appPort, isAppPortSet := os.LookupEnv("DCS_LISTEN_PORT")

	if !isAppPortSet {
		println("DCS_LISTEN_PORT not set, defaulting to 8080")
		appPort = "8080"
	}

	AppPort = appPort

	dbUser, isDbUserSet := os.LookupEnv("DB_USERNAME")
	dbPass, isDbPassSet := os.LookupEnv("DB_PASSWORD")
	dbName, isDbNameSet := os.LookupEnv("DB_NAME")
	dbPort, isDbPortSet := os.LookupEnv("DB_PORT")
	dbHost, isDbHostSet := os.LookupEnv("DB_HOST")

	if !isDbPassSet {
		panic("DB_PASSWORD environment variable not set")
	}

	if !isDbUserSet {
		panic("DB_USERNAME environment variable not set")
	}

	if !isDbNameSet {
		println("DB_NAME environment variable not set, defaulting to dcs")
		dbName = "dcs"
	}

	if !isDbPortSet {
		println("DB_PORT environment variable not set, defaulting to 3306")
		dbPort = "3306"
	}

	if !isDbHostSet {
		println("DB_HOST environment variable not set, defaulting to 127.0.0.1")
		dbHost = "127.0.0.1"
	}

	DbUser = dbUser
	DbPass = dbPass
	DbName = dbName
	DbPort = dbPort
	DbHost = dbHost

}
