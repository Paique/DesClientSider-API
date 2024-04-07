package main

import (
	"database/sql"
	"dcs-rest-api/data"
	"dcs-rest-api/util"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

func CreateDbInstance() {
	data.ConnTries++
	dsn := data.DbUser + ":" + data.DbPass + "@tcp(" + data.DbHost + ":" + data.DbPort + ")/" + data.DbName

	db, err := sql.Open("mysql", dsn)

	//Todo: I have no idea if this is right
	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(0)

	if err != nil {
		panic(err)
	}

	data.DbConn = db

	if data.ConnTries > 10 {
		panic("Cannot connect to database after 10 retries")
	}

	err = db.Ping()

	if err != nil {
		println("Cannot connect to database!")
		println("Trying again in 5 secs...")
		time.Sleep(5 * time.Second)
		println("Try: " + strconv.Itoa(data.ConnTries) + "/10")
		CreateDbInstance()
	}

	println("Database connection established!")

	println("Initializing database...")
	initTables()
}

func initTables() {
	db := data.DbConn

	_, tableCheck := db.Query("select * from Keywords")

	var err error
	if tableCheck != nil {
		println("Creating table Keywords")
		_, err = db.Query("CREATE TABLE Keywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")
		keywords := util.GetdefaultList()

		println("Adding default keywords to the table")
		for _, element := range keywords {
			println("Added " + element + " to the table keywords")
			_, err = db.Query("INSERT INTO Keywords(keyword) VALUES (?);", element)
		}
	}

	_, tableCheck = db.Query("select * from ContraKeywords")

	if tableCheck != nil {
		_, err = db.Query("CREATE TABLE ContraKeywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")

	}

	if err != nil {
		println("Cannot create tables")
		panic(err)
	}
}

func GetKeysList() []data.Keywords {
	var keywords []data.Keywords

	db := data.DbConn
	if db == nil {
		panic("Database connection failed")
	}
	rows, err := db.Query("select * from Keywords")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	for rows.Next() {
		var dbModkey string
		err = rows.Scan(&dbModkey)
		if err != nil {
			panic(err)
		}
		keywords = append(keywords, data.Keywords{ID: strconv.Itoa(id), Keyword: dbModkey})
		id++
	}
	return keywords
}

func GetContraKeyList() []data.ContraKeys {
	db := data.DbConn
	rows, err := db.Query("select * from ContraKeywords")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var contraKey []data.ContraKeys
	for rows.Next() {
		var dbModContrakey string
		err = rows.Scan(&dbModContrakey)
		if err != nil {
			panic(err)
		}
		contraKey = append(contraKey, data.ContraKeys{ID: strconv.Itoa(id), Keyword: dbModContrakey})
		id++
	}
	return contraKey
}
