// Package postgres
// Provides functions to access PostgreSQL database
package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"shopping-list/pkg/item"
	"shopping-list/pkg/logging"
)

var conn *pgxpool.Pool

// CreateConnection create database *connection
func CreateConnection() {
	logging.LogInfo("Attempting to *connect to database")

	var err error
	//conn, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	conn, err = pgxpool.New(context.Background(), "postgresql://test:asdasd@172.22.0.2:5432/shopping")
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

// InsertItem inserts a new item into the database.
func InsertItem(name string, note string, amount int) {
	query := fmt.Sprintf("INSERT INTO items (name, note, amount, status) VALUES ('%s','%s','%d','new');", name, note, amount)
	executeQuery(query)
}

// CreateTable create new table for given channel
// runs at start to make sure tables exists
func CreateTable() {
	query := "create table if not exists items (" +
		"id SERIAL PRIMARY KEY," +
		"name VARCHAR(100) NOT NULL," +
		"note VARCHAR(100) NOT NULL," +
		"amount INTEGER," +
		"status VARCHAR(10)" +
		");"

	executeQuery(query)
}

// executeQuery executes a query using a *connection from the pool
func executeQuery(query string) {
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		logging.LogError("Failed to execute query", err)
	} else {
		queryMessage := fmt.Sprintf("executed query successful!\n\t%s\n", query)
		logging.LogInfo(queryMessage)
	}
}

// GetItem returns single item by given ID.
func GetItem(id string) item.Item {
	query := fmt.Sprintf("select * from items where id = %s", id)
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		logging.LogWarning("Failed to find item with id " + id)
	}

	var i item.Item
	for rows.Next() {
		err := rows.Scan(&i.ID, &i.Name, &i.Note, &i.Amount, &i.Status)
		if err != nil {
			logging.LogWarning("Failed to scann row: " + err.Error())
		}
	}

	return i
}

// UpdateItemStatus changes status of an item from old to new and other way around.
func UpdateItemStatus(id string) {
	i := GetItem(id)

	if i.Status == "new" {
		query := fmt.Sprintf("UPDATE items SET status = 'old' WHERE id = '%s'", id)
		executeQuery(query)
	} else if i.Status == "old" {
		query := fmt.Sprintf("UPDATE items SET status = 'new' WHERE id = '%s'", id)
		executeQuery(query)
	}
}

func ChangeItem(id, name, note string, amount int) {
	query := fmt.Sprintf("UPDATE items SET name = '%s', note = '%s', amount = '%d' WHERE id = '%s'", name, note, amount, id)
	executeQuery(query)
}

// DeleteItemStatus changes the status of an item from new or old to deleted and from deleted to old.
// Used on managing page for removing items.
func DeleteItemStatus(id string) {
	i := GetItem(id)

	if i.Status == "new" || i.Status == "old" {
		query := fmt.Sprintf("UPDATE items SET status = 'deleted' WHERE id = '%s'", id)
		executeQuery(query)
	} else if i.Status == "deleted" {
		query := fmt.Sprintf("UPDATE items SET status = 'old' WHERE id = '%s'", id)
		executeQuery(query)
	}
}

// GetItems gets a slice containing all items that match the given status.
func GetItems(status string) []item.Item {
	// use LIKE here to be able to match using wildcards
	query := fmt.Sprintf("SELECT * FROM items WHERE status LIKE '%s'", status)
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		logging.LogWarning("Failed to query for items with status " + status)
	}

	var itemList []item.Item
	for rows.Next() {
		var i item.Item
		// scan row into item
		err := rows.Scan(&i.ID, &i.Name, &i.Note, &i.Amount, &i.Status)
		if err != nil {
			logging.LogWarning("Failed to scan row: " + err.Error())
		}

		itemList = append(itemList, i)
	}

	return itemList
}
