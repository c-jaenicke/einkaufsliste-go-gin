// Package postgres
// Provides functions to access PostgreSQL database
package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"shopping-list/pkg/logging"
)

var conn *pgxpool.Pool

// CreateConnection create database *connection
func CreateConnection() {
	logging.LogInfo("Attempting to *connect to database")

	var err error
	// conn, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	// Line for testing locally on test db
	conn, err = pgxpool.New(context.Background(), "postgresql://test:asdasd@172.21.0.2:5432/shopping")
	if err != nil {
		logging.LogPanic("Failed to *connect to database", err)
		os.Exit(1)
	}

	pingErr := conn.Ping(context.Background())
	if pingErr != nil {
		logging.LogPanic("Initial ping to database failed", err)
		os.Exit(1)
	}

	logging.LogInfo("Connection to database successful")
}

// PingDatabase sends ping to database to test if connection is still up.
// Returns true if connection is still up.
func PingDatabase() bool {
	pingErr := conn.Ping(context.Background())
	if pingErr != nil {
		logging.LogError("Ping to database failed", pingErr)
		return false
	} else {
		logging.LogInfo("Ping to database successful")
		return true
	}
}

// CreateTable creates new tables. Includes single value for category table.
func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS category
	(
		id   SERIAL PRIMARY KEY,
		name VARCHAR(100),
		color CHAR(7) DEFAULT '#FFFFFF',
		color_name VARCHAR(30) DEFAULT 'WHITE'
	);`
	executeStatement(query)

	query = `INSERT INTO category (id, name) VALUES (0, 'Keine');`
	executeStatement(query)

	query = `create table if not exists items
	(
    	id     SERIAL PRIMARY KEY,
    	name   VARCHAR(100) NOT NULL,
    	note   VARCHAR(100) NOT NULL,
    	amount INTEGER,
    	status VARCHAR(10) DEFAULT 'new',
    	cat_id INT DEFAULT '0',
    	CONSTRAINT fk_cat_id
        	FOREIGN KEY (cat_id)
            	REFERENCES category(id)
            	ON DELETE SET DEFAULT
	);`
	executeStatement(query)

}

// executeStatement executes a statement on a table
func executeStatement(query string) {
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		logging.LogError("Failed to execute statement: "+query, err)
	} else {
		queryMessage := fmt.Sprintf("executed query successful!\n\t%s\n", query)
		logging.LogInfo(queryMessage)
	}
}

// executeQuery executes a given query and returns the rows.
func executeQuery(query string) pgx.Rows {
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		logging.LogError("Failed to execute query: "+query, err)
	}

	return rows
}
