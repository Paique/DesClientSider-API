package main

import (
	"database/sql"
	"dcs-api/data"
	"dcs-api/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"time"
)

var conn *sql.DB

// GetDbInstance returns the database connection instance
func GetDbInstance() *sql.DB {

	if conn == nil {
		connectDB()
	}

	return conn
}

var try = 0

// connectDB creates a new database connection instance
func connectDB() {

	if try >= 10 {
		log.Panic("Can't connect to database after 10 attempts")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", data.DbUser, data.DbPass, data.DbHost, data.DbPort, data.DbName)

	var err error
	conn, err = sql.Open("mysql", dsn)
	conn.SetMaxOpenConns(10)

	if conn.Ping() != nil || err != nil {
		try++
		log.Printf("Cannot ping database, trying to connect again in 5 secs.")
		time.Sleep(5 * time.Second)
		fmt.Printf("Trying to connect to database %d/10 \n", try)
		connectDB()
		return
	}

	initTables()

	log.Println("database connection initialized")
}

func initTables() {
	var err error

	db := GetDbInstance()

	rows, err := db.Query("select * from Keywords")

	if err != nil {
		log.Println("Creating table Keywords")
		rows, err = db.Query("CREATE TABLE Keywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")

		if err != nil {
			log.Panic("Cannot create table Keywords: ", err)
		}

		keywords := util.GetdefaultKeys(false)

		rows.Close()

		println("Adding default keywords to the table")
		for _, element := range keywords {
			rows, err = db.Query("INSERT INTO Keywords(keyword) VALUES (?);", element)

			if err != nil {
				log.Printf("Failed inserting keyword %s into table", element)
			} else {
				log.Println("Added " + element + " to the table keywords")
			}

			rows.Close()
		}
		println("Table keywords: Defaults generated")
	} else {
		println("Table keywords already exists")
		rows.Close()
	}

	rows, err = db.Query("select * from ContraKeywords")

	if err != nil {
		rows, err = db.Query("CREATE TABLE ContraKeywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")
		rows.Close()

		if err != nil {
			log.Panic("Cannot create table Keywords: ", err)
		}

		keywords := util.GetdefaultKeys(true)

		for _, element := range keywords {
			rows, err = db.Query("INSERT INTO ContraKeywords(keyword) VALUES (?);", element)
			rows.Close()
			if err != nil {
				log.Printf("Failed inserting keyword %s into table", element)
			} else {
				log.Println("Added " + element + " to the table contraKeywords")
			}
		}
		println("Table contaKeywords: Defaults generated")
	} else {
		println("Table contaKeywords already exists")
		rows.Close()
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
