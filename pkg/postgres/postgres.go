// Package postgres
// Provides functions to access PostgreSQL database
package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"shopping-list/pkg/category"
	"shopping-list/pkg/item"
	"shopping-list/pkg/logging"
)

var conn *pgxpool.Pool

//
// DATABASE FUNCTIONS
//

// CreateConnection create database *connection
func CreateConnection() {
	logging.LogInfo("Attempting to *connect to database")

	var err error
	// TODO change to env again here
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

// CreateTable create new table for given channel
// runs at start to make sure tables exists
func CreateTable() {
	// create table containing items
	query := "create table if not exists items (" +
		"id SERIAL PRIMARY KEY," +
		"name VARCHAR(100) NOT NULL," +
		"note VARCHAR(100) NOT NULL," +
		"amount INTEGER," +
		"status VARCHAR(10)" +
		");"

	executeQuery(query)

	// TODO add query to include category table
	// TODO add query to insert category "0 - None"
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

//
// ITEM FUNCTIONS
//

// InsertItem inserts a new item into the database.
func InsertItem(name string, note string, amount int, cat_id string) {
	query := fmt.Sprintf("INSERT INTO items (name, note, amount, status, cat_id) VALUES ('%s','%s','%d','new','%s');", name, note, amount, cat_id)
	executeQuery(query)
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
		err := rows.Scan(&i.ID, &i.Name, &i.Note, &i.Amount, &i.Status, &i.Cat_id)
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
		err := rows.Scan(&i.ID, &i.Name, &i.Note, &i.Amount, &i.Status, &i.Cat_id)
		if err != nil {
			logging.LogWarning("Failed to scan row: " + err.Error())
		}

		itemList = append(itemList, i)
	}

	return itemList
}

//
// CATEGORY FUNCTIONS
//

// CreateCategory creates a new category with the given name
func CreateCategory(name string) {
	query := fmt.Sprintf("INSERT INTO category (name) values ('%s');", name)
	executeQuery(query)
}

// GetAllCategories gets slice of all categories that exist
func GetAllCategories() []category.Category {
	query := fmt.Sprintf("SELECT * FROM category;")
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		logging.LogWarning("Failed to query all categories " + query)
	}

	var catList []category.Category
	for rows.Next() {
		var c category.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			logging.LogWarning("Failed to scan row: " + err.Error())
		}

		catList = append(catList, c)
	}

	return catList
}

func GetCategory(id string) category.Category {
	query := fmt.Sprintf("SELECT * FROM category WHERE id = '%s';", id)
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		logging.LogWarning("Failed to query all categories " + query)
	}

	var c category.Category
	for rows.Next() {
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			logging.LogWarning("Failed to scan row: " + err.Error())
		}
	}

	return c
}

func ChangeCategory(id, name string) {
	query := fmt.Sprintf("UPDATE category SET name = '%s' WHERE id = '%s'", name, id)
	executeQuery(query)
}

func GetItemsInCategory(id string, status string) []item.Item {
	query := fmt.Sprintf("SELECT * FROM items WHERE cat_id = '%s' AND status LIKE '%s';", id, status)

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		logging.LogWarning("Failed query" + query)
	}

	var itemList []item.Item
	for rows.Next() {
		var i item.Item
		// scan row into item
		err := rows.Scan(&i.ID, &i.Name, &i.Note, &i.Amount, &i.Status, &i.Cat_id)
		if err != nil {
			logging.LogWarning("Failed to scan row: " + err.Error())
		}

		itemList = append(itemList, i)
	}

	return itemList
}

// DeleteCategory deletes an existing category from the table
func DeleteCategory(id string) {
	query := fmt.Sprintf("DELETE FROM category WHERE id = '%s';", id)
	executeQuery(query)
}
