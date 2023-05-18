package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"shopping-list/pkg/item"
	"shopping-list/pkg/logging"
)

// SaveItem save a single item
func SaveItem(item item.Item) {
	query := fmt.Sprintf("INSERT INTO items (name, note, amount, status, cat_id) VALUES ('%s','%s','%d','new','%d');", item.Name, item.Note, item.Amount, item.Cat_id)
	executeStatement(query)
}

// UpdateItem update a single item
func UpdateItem(item item.Item, id int) {
	query := fmt.Sprintf("UPDATE items SET name = '%s', note = '%s', amount = '%d', cat_id = '%d' WHERE id = '%d';", item.Name, item.Note, item.Amount, item.Cat_id, id)
	executeStatement(query)
}

// UpdateItemStatus change the status of the item with the given id, from 'old' to 'new'
func SwitchItemStatus(id int) {
	i := GetItem(id)

	if i.Status == "new" {
		query := fmt.Sprintf("UPDATE items SET status = 'old' WHERE id = '%d';", id)
		executeStatement(query)
	} else if i.Status == "old" {
		query := fmt.Sprintf("UPDATE items SET status = 'new' WHERE id = '%d';", id)
		executeStatement(query)
	}
}

// DeleteItem deletes item with given id
func DeleteItem(id int) {
	query := fmt.Sprintf("DELETE FROM items WHERE id = '%d';", id)
	executeStatement(query)
}

// DeleteAllItems EMPTYS WHOLE TABLE
func DeleteAllItems() {
	query := fmt.Sprintf("TRUNCATE items;")
	executeStatement(query)
}

// GetItem get a single item object with the id
func GetItem(id int) item.Item {
	query := fmt.Sprintf("SELECT * FROM items WHERE id = '%d';", id)
	rows := executeQuery(query)

	return rowsToItems(rows)[0]
}

// GetAllItems get all items in the table
func GetAllItems() []item.Item {
	query := fmt.Sprintf("SELECT * FROM items;")
	rows := executeQuery(query)

	return rowsToItems(rows)
}

// GetItemsWithStatus get all items with the given status. Can take '%' as wildcard for all statuses
func GetItemsWithStatus(status string) []item.Item {
	query := fmt.Sprintf("SELECT * FORM items WHERE status = '%s';", status)
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
