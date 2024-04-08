package main

import (
	"database/sql"
	"dcs-rest-api/data"
	"dcs-rest-api/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

var conn *sql.DB

// GetDbInstance returns the database connection instance
func GetDbInstance() *sql.DB {

	if conn == nil {
		connectDB()
	}

	return conn
}

// CreateDbInstance creates a new database connection instance
func connectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", data.DbUser, data.DbPass, data.DbHost, data.DbPort, data.DbName)

	var err error
	conn, err = sql.Open("mysql", dsn)
	conn.SetMaxOpenConns(10)

	if err != nil {
		log.Panicf("cannot connect to database: %s", err)
		return
	}

	initTables()

	log.Println("database connection initialized")
}

func initTables() {
	db := GetDbInstance()

	_, tableCheck := db.Query("select * from Keywords")

	var err error
	if tableCheck != nil {
		log.Println("Creating table Keywords")
		_, err = db.Query("CREATE TABLE Keywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")
		keywords := util.GetdefaultList()

		println("Adding default keywords to the table")
		for _, element := range keywords {
			log.Println("Added " + element + " to the table keywords")
			_, err = db.Query("INSERT INTO Keywords(keyword) VALUES (?);", element)
		}
	}

	_, tableCheck = db.Query("select * from ContraKeywords")

	if tableCheck != nil {
		_, err = db.Query("CREATE TABLE ContraKeywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")

	}

	if err != nil {
		log.Println("Cannot create tables")
		log.Panicln(err)
	}
}

func GetKeysList() []data.Keywords {
	db := GetDbInstance()

	rows, err := db.Query("SELECT * FROM Keywords")
	if err != nil {
		log.Panicf("cannot query database: %s", err)
		return nil
	}
	defer rows.Close()

	var id int
	var keywords []data.Keywords

	for rows.Next() {
		var dbModkey string

		err = rows.Scan(&dbModkey)
		if err != nil {
			log.Panicf("cannot scan row: %s", err)
			return nil
		}

		keywords = append(keywords, data.Keywords{
			ID:      strconv.Itoa(id),
			Keyword: dbModkey,
		})

		id++
	}

	return keywords
}

func GetContraKeyList() []data.ContraKeys {
	db := GetDbInstance()

	rows, err := db.Query("SELECT * FROM ContraKeywords")
	if err != nil {
		log.Panicf("cannot query database: %s", err)
		return nil
	}
	defer rows.Close()

	var id int
	var contraKey []data.ContraKeys

	for rows.Next() {
		var dbModContrakey string

		err = rows.Scan(&dbModContrakey)
		if err != nil {
			log.Panicf("cannot scan row: %s", err)
		}

		contraKey = append(contraKey, data.ContraKeys{
			ID:      strconv.Itoa(id),
			Keyword: dbModContrakey,
		})

		id++
	}

	return contraKey
}
