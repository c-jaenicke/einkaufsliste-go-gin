// Package postgres
// Provides functions to access PostgreSQL database
package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
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
    	status VARCHAR(10),
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

//
// ITEM FUNCTIONS
//

// InsertItem inserts a new item into the database.
func InsertItem(name string, note string, amount int, cat_id string) {
	query := fmt.Sprintf("INSERT INTO items (name, note, amount, status, cat_id) VALUES ('%s','%s','%d','new','%s');", name, note, amount, cat_id)
	executeStatement(query)
}

// GetItem returns single item by given ID.
func GetItem(id string) item.Item {
	query := fmt.Sprintf("select * from items where id = %s", id)
	rows := executeQuery(query)

	return rowsToItems(rows)[0]
}

// UpdateItemStatus changes status of an item from old to new and other way around.
func UpdateItemStatus(id string) {
	i := GetItem(id)

	if i.Status == "new" {
		query := fmt.Sprintf("UPDATE items SET status = 'old' WHERE id = '%s'", id)
		executeStatement(query)
	} else if i.Status == "old" {
		query := fmt.Sprintf("UPDATE items SET status = 'new' WHERE id = '%s'", id)
		executeStatement(query)
	}
}

func ChangeItem(id string, name string, note string, amount int, cat_id string) {
	query := fmt.Sprintf("UPDATE items SET name = '%s', note = '%s', amount = '%d', cat_id = '%s' WHERE id = '%s';", name, note, amount, cat_id, id)
	executeStatement(query)
}

// DeleteItemStatus changes the status of an item from new or old to deleted and from deleted to old.
// Used on managing page for removing items.
func DeleteItemStatus(id string) {
	i := GetItem(id)

	if i.Status == "new" || i.Status == "old" {
		query := fmt.Sprintf("UPDATE items SET status = 'deleted' WHERE id = '%s'", id)
		executeStatement(query)
	} else if i.Status == "deleted" {
		query := fmt.Sprintf("UPDATE items SET status = 'old' WHERE id = '%s'", id)
		executeStatement(query)
	}
}

// GetItems gets a slice containing all items that match the given status.
func GetItems(status string) []item.Item {
	// use LIKE here to be able to match using wildcards
	query := fmt.Sprintf("SELECT * FROM items WHERE status LIKE '%s'", status)
	rows := executeQuery(query)

	return rowsToItems(rows)
}

// rowsToItems turn query response containing items to slice of items
func rowsToItems(rows pgx.Rows) []item.Item {
	var itemSlice []item.Item
	for rows.Next() {
		var i item.Item
		err := rows.Scan(&i.ID, &i.Name, &i.Note, &i.Amount, &i.Status, &i.Cat_id)
		if err != nil {
			logging.LogWarning("Failed to scan row into item object: " + err.Error())
		}

		itemSlice = append(itemSlice, i)
	}

	return itemSlice
}

// DeleteAllItems permanently deletes all entries in the items table. NOT REVERSIBLE
func DeleteAllItems() {
	query := fmt.Sprintf("TRUNCATE items;")
	executeStatement(query)
}

//
// CATEGORY FUNCTIONS
//

// CreateCategory creates a new category with the given name
func CreateCategory(name string, color string, colorName string) {
	query := fmt.Sprintf("INSERT INTO category (name, color, color_name) values ('%s', '%s', '%s');", name, color, colorName)
	executeStatement(query)
}

// GetAllCategories gets slice of all categories that exist
func GetAllCategories() []category.Category {
	query := fmt.Sprintf("SELECT * FROM category;")
	rows := executeQuery(query)

	return rowsToCategory(rows)
}

// GetCategory gets category object from given id
func GetCategory(id string) category.Category {
	query := fmt.Sprintf("SELECT * FROM category WHERE id = '%s';", id)
	rows := executeQuery(query)

	return rowsToCategory(rows)[0]
}

// ChangeCategory updates name of a category
func ChangeCategory(id, name string) {
	query := fmt.Sprintf("UPDATE category SET name = '%s' WHERE id = '%s'", name, id)
	executeStatement(query)
}

// GetItemsInCategory returns slice of all items in a category with the given status
func GetItemsInCategory(id string, status string) []item.Item {
	query := fmt.Sprintf("SELECT * FROM items WHERE cat_id = '%s' AND status LIKE '%s';", id, status)
	rows := executeQuery(query)

	return rowsToItems(rows)
}

// DeleteCategory deletes an existing category from the table
func DeleteCategory(id string) {
	query := fmt.Sprintf("DELETE FROM category WHERE id = '%s';", id)
	executeStatement(query)
}

// rowsToCategory transforms rows into category objects. Returned as a slice of category
func rowsToCategory(rows pgx.Rows) []category.Category {
	var categorySlice []category.Category
	for rows.Next() {
		var c category.Category
		err := rows.Scan(&c.ID, &c.Name, &c.Color, &c.Color_name)
		if err != nil {
			logging.LogWarning("Failed to scan row into category object: " + err.Error())
		}

		categorySlice = append(categorySlice, c)
	}

	return categorySlice
}
