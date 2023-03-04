package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"shopping-list/pkg/item"
	"shopping-list/pkg/logging"
)

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
