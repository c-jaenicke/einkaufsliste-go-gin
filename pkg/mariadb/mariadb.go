// Package mariadb
// not used anymore, got replaced by postgres
package mariadb

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

// create database handle
func CreateConnection() sql.DB {
	fmt.Fprint(os.Stdout, "[Info] Connecting to db...\n")
	var db *sql.DB
	var cfg = mysql.Config{
		User:   "root",
		Passwd: "secret",
		Net:    "tcp",
		//Addr:                 "mariadb:3306",
		Addr:                 os.Getenv("DB_ADDRESS") + ":3306",
		DBName:               "links",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] connecting to database!\n\t%s\n", err.Error())
		os.Exit(1)
	}

	// set connection limits
	db.SetConnMaxLifetime(time.Hour * 24)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// confirm that db accepts connection
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Fprintf(os.Stderr, "[Error] initial ping to db failed!\n\t%s\n", pingErr.Error())
		os.Exit(1)
	}

	fmt.Fprint(os.Stdout, "[Info] Connection successful!\n")

	return *db
}

func PingDatabase(db sql.DB) bool {
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Fprint(os.Stderr, "[Error] lost connection to database!")
		return false
	} else {
		return true
	}
}

// insert slice of links
func InsertLinks(links []string, db sql.DB, channel string) {
	fmt.Fprintf(os.Stdout, "[Info] inserting %d links from channel %s\n", len(links), channel)

	// build query string
	query := fmt.Sprintf("INSERT INTO %s (url) VALUES ", channel)
	var sb strings.Builder
	sb.WriteString(query)
	for i, value := range links {
		if i == len(links)-1 {
			sb.WriteString(fmt.Sprintf("(\"%s\")", value))
		} else {
			sb.WriteString(fmt.Sprintf("(\"%s\"), ", value))
		}
	}
	sb.WriteString(";")

	// execute insertion query
	executeQuery(sb.String(), db)
}

// create new table for a channel if it doesn't already exist
func CreateTable(db sql.DB, channel string) {
	query := "USE links;"
	executeQuery(query, db)

	query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id INT AUTO_INCREMENT NOT NULL, time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, url VARCHAR(250) NOT NULL, PRIMARY KEY (id));", channel)
	executeQuery(query, db)
}

func executeQuery(query string, db sql.DB) {
	_, err := db.Exec(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] error performing query!\n\t%s\n\t%s\n", err.Error(), query)
	} else {
		fmt.Fprintf(os.Stdout, "[Info] query successful!\n\t%s\n", query)
	}
}
