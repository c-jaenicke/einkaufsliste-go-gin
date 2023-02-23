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

// CreateConnection create database connection
func CreateConnection() pgxpool.Pool {
	logging.LogInfo("Attempting to connect to database")

	//conn, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	conn, err := pgxpool.New(context.Background(), "postgresql://test:asdasd@172.22.0.2:5432/shopping")
	if err != nil {
		logging.LogPanic("Failed to connect to database", err)
		os.Exit(1)
	}

	pingErr := conn.Ping(context.Background())
	if pingErr != nil {
		logging.LogPanic("Initial ping to database failed", err)
		os.Exit(1)
	}

	logging.LogInfo("Connection to database successful")

	return *conn
}

// PingDatabase send ping to database
func PingDatabase(conn pgxpool.Pool) bool {
	pingErr := conn.Ping(context.Background())
	if pingErr != nil {
		logging.LogError("Ping to database failed", pingErr)
		return false
	} else {
		logging.LogInfo("Ping to database successful")
		return true
	}
}

func InsertItem(name string, note string, amount int, conn pgxpool.Pool) {
	message := fmt.Sprintf("Inserting item ", name, note, amount)
	logging.LogInfo(message)

	var item = item.Item{Name: name,
		Note:   note,
		Amount: amount,
		Status: "new",
	}

	query := fmt.Sprintf("INSERT INTO items (name, note, amount, status) VALUES ("+
		"'%s',"+
		"'%s',"+
		"'%d',"+
		"'%s'"+
		");", item.Name, item.Note, item.Amount, item.Status)

	logging.LogInfo(query)
	executeQuery(query, conn)
}

// CreateTable create new table for given channel
// runs at start to make sure tables exists
func CreateTable(conn pgxpool.Pool) {
	query := "create table if not exists items (" +
		"id SERIAL PRIMARY KEY," +
		"name VARCHAR(100) NOT NULL," +
		"note VARCHAR(100) NOT NULL," +
		"amount INTEGER," +
		"status VARCHAR(10)" +
		");"

	executeQuery(query, conn)
}

// executeQuery executes a query using a connection from the pool
func executeQuery(query string, conn pgxpool.Pool) {
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		logging.LogError("Failed to execute query", err)
	} else {
		queryMessage := fmt.Sprintf("executed query successful!\n\t%s\n", query)
		logging.LogInfo(queryMessage)
	}
}

func GetAllItems(conn pgxpool.Pool) []item.Item {
	query := "select * from items"
	rows, err := conn.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		logging.LogWarning("Failed to query for items")
	}

	var items []item.Item
	for rows.Next() {
		var item item.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Note, &item.Amount, &item.Status)
		if err != nil {
			logging.LogWarning("Failed to scann row: " + err.Error())
		}

		items = append(items, item)
	}

	return items
}

func GetNewItems(conn pgxpool.Pool) []item.Item {
	query := "select * from items where status = 'new'"
	rows, err := conn.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		logging.LogWarning("Failed to query for items")
	}

	var items []item.Item
	for rows.Next() {
		var item item.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Note, &item.Amount, &item.Status)
		if err != nil {
			logging.LogWarning("Failed to scann row: " + err.Error())
		}

		items = append(items, item)
	}

	return items
}

func GetOldItems(conn pgxpool.Pool) []item.Item {
	query := "select * from items where status = 'old'"
	rows, err := conn.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		logging.LogWarning("Failed to query for items")
	}

	var items []item.Item
	for rows.Next() {
		var item item.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Note, &item.Amount, &item.Status)
		if err != nil {
			logging.LogWarning("Failed to scann row: " + err.Error())
		}

		items = append(items, item)
	}

	return items
}

func GetItem(id string, conn pgxpool.Pool) item.Item {
	query := fmt.Sprintf("select * from items where id = %s", id)
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		logging.LogWarning("Failed to find item with id " + id)
	}

	var item item.Item
	for rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.Note, &item.Amount, &item.Status)
		if err != nil {
			logging.LogWarning("Failed to scann row: " + err.Error())
		}
	}

	return item
}

func UpdateItemStatus(id string, conn pgxpool.Pool) {
	item := GetItem(id, conn)

	if item.Status == "new" {
		query := fmt.Sprintf("UPDATE items SET status = 'old' WHERE id = '%s'", id)
		executeQuery(query, conn)
	} else if item.Status == "old" {
		query := fmt.Sprintf("UPDATE items SET status = 'new' WHERE id = '%s'", id)
		executeQuery(query, conn)
	}
}

func DeleteItemStatus(id string, conn pgxpool.Pool) {
	item := GetItem(id, conn)

	if item.Status == "new" || item.Status == "old" {
		query := fmt.Sprintf("UPDATE items SET status = 'deleted' WHERE id = '%s'", id)
		executeQuery(query, conn)
	} else if item.Status == "deleted" {
		query := fmt.Sprintf("UPDATE items SET status = 'old' WHERE id = '%s'", id)
		executeQuery(query, conn)
	}
}

func GetDeletedItems(conn pgxpool.Pool) []item.Item {
	query := "select * from items where status = 'deleted'"
	rows, err := conn.Query(context.Background(), query)
	defer rows.Close()
	if err != nil {
		logging.LogWarning("Failed to query for items")
	}

	var items []item.Item
	for rows.Next() {
		var item item.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Note, &item.Amount, &item.Status)
		if err != nil {
			logging.LogWarning("Failed to scann row: " + err.Error())
		}

		items = append(items, item)
	}

	return items
}
