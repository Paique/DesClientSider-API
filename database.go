package main

import (
	"database/sql"
	"dcs-rest-api/data"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func accessDb() *sql.DB {
	dsn := data.DbUser + ":" + data.DbPass + "@tcp(127.0.0.1:" + data.DbPort + ")/" + data.DbName

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return db
}

func GetKeysList() []data.Keywords {
	db := accessDb()
	rows, err := db.Query("select * from Keywords")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var keywords []data.Keywords
	for rows.Next() {
		var dbModkey string
		err = rows.Scan(&dbModkey)
		if err != nil {
			panic(err)
		}
		keywords = append(keywords, data.Keywords{ID: strconv.Itoa(id), Keyword: dbModkey})
		id++
	}
	defer db.Close()
	return keywords
}

func GetContraKeyList() []data.ContraKeys {
	db := accessDb()
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
	defer db.Close()
	return contraKey
}
