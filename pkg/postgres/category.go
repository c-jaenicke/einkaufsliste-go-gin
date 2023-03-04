package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"shopping-list/pkg/category"
	"shopping-list/pkg/item"
	"shopping-list/pkg/logging"
)

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
