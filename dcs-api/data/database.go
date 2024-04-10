package data

import (
	"database/sql"
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPass, DbHost, DbPort, DbName)

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

		log.Println("Adding default keywords to the table")
		for _, element := range keywords {
			rows, err = db.Query("INSERT INTO Keywords(keyword) VALUES (?);", element)

			if err != nil {
				log.Printf("Failed inserting keyword %s into table", element)
			} else {
				log.Println("Added " + element + " to the table Keywords")
			}

			rows.Close()
		}
		log.Println("Table Keywords: Defaults generated")
	} else {
		log.Println("Table Keywords already exists")
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
		log.Println("Table ContaKeywords: Defaults generated")
	} else {
		log.Println("Table ContaKeywords already exists")
		rows.Close()
	}

	rows, err = db.Query("select * from Logs")
	if err != nil {
		rows, err = db.Query("CREATE TABLE Logs(time DATETIME, ip varchar(20), url VARCHAR(20));")
		rows.Close()

		if err != nil {
			log.Panic("Cannot create table Logs: ", err)
		}

	} else {
		log.Println("Table Logs already exists")
	}

}

func GetKeysList() []Keywords {
	db := GetDbInstance()

	rows, err := db.Query("SELECT * FROM Keywords")
	if err != nil {
		log.Panicf("Cannot query database: %s", err)
		return nil
	}
	defer rows.Close()

	var id int
	var keywords []Keywords

	for rows.Next() {
		var dbModkey string

		err = rows.Scan(&dbModkey)
		if err != nil {
			log.Panicf("Cannot scan row: %s", err)
			return nil
		}

		keywords = append(keywords, Keywords{
			ID:      strconv.Itoa(id),
			Keyword: dbModkey,
		})

		id++
	}

	return keywords
}

func GetContraKeyList() []ContraKeys {
	db := GetDbInstance()

	rows, err := db.Query("SELECT * FROM ContraKeywords")
	if err != nil {
		log.Panicf("Cannot query database: %s", err)
		return nil
	}
	defer rows.Close()

	var id int
	var contraKey []ContraKeys

	for rows.Next() {
		var dbModContrakey string

		err = rows.Scan(&dbModContrakey)
		if err != nil {
			log.Panicf("Cannot scan row: %s", err)
		}

		contraKey = append(contraKey, ContraKeys{
			ID:      strconv.Itoa(id),
			Keyword: dbModContrakey,
		})

		id++
	}

	return contraKey
}
