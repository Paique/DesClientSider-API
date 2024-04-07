package main

import (
	"database/sql"
	"dcs-rest-api/data"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"sync"
)

var conn *sql.DB
var lock sync.Mutex

// GetDbInstance returns the database connection instance
func GetDbInstance() *sql.DB {
	if conn == nil {
		lock.Lock()
		defer lock.Unlock()

		if conn == nil {
			connectDB()
		}
	}

	return conn
}

// CreateDbInstance creates a new database connection instance
func connectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s", data.DbUser, data.DbPass, data.DbPort, "dcs")

	conn, err := sql.Open("mysql", dsn)
	conn.SetMaxOpenConns(10)
	if err != nil {
		log.Panicf("cannot connect to database: %s", err)
		return
	}

	conn = conn

	initTables()

	log.Println("database connection initialized")
}

func initTables() {
	db := GetDbInstance()
	_, err := db.Query("CREATE TABLE IF NOT EXISTS Keywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")
	if err != nil {
		log.Panicln("cannot create table Keywords")
		return
	}

	_, err = db.Query("CREATE TABLE IF NOT EXISTS ContraKeywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")
	if err != nil {
		log.Panicln("cannot create table ContraKeywords")
		return
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
