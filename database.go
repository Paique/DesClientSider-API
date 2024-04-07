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

	initTables()
}

func initTables() {
	db := data.DbConn
	_, err := db.Query("CREATE TABLE IF NOT EXISTS Keywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")

	if err != nil {
		panic("Cannot create table Keywords")
		return
	}
	_, err = db.Query("CREATE TABLE IF NOT EXISTS ContraKeywords(keyword VARCHAR(30) PRIMARY KEY UNIQUE);")

	if err != nil {
		panic("Cannot create table ContraKeywords")
		return
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
