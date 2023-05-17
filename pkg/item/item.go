// Package item
// Contains struct which embodies an item
package item

type Item struct {
	// TODO possibly expand query to include default values for category = 0 and status = new
	ID     int    `json:"id"`     // id of an item
	Name   string `json:"name"`   // main description of an item
	Note   string `json:"note"`   // note on the item itself
	Amount int    `json:"amount"` // how much of an item should be bought
	Status string `json:"status"` // status of item, can be [new old deleted] can be expanded throughout development
	Cat_id int    `json:"cat_Id"` // foreign key, id of category, snake_case to keep consistent with database schema
}
