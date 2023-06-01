package postgres

import (
	"shopping-list/pkg/item"
	"testing"
)

var i = item.Item{
	Name:   "TEST ITEM",
	Note:   "TEST ITEM",
	Amount: 1,
	Status: "new",
	Store:  "edeka",
	Cat_id: 0,
}

func TestMain(m *testing.M) {
	prepDatabase()
	m.Run()
	DeleteAllItems()
}

func prepDatabase() {
	CreateConnection()
	// uncomment this function if tables dont exist yet
	//CreateTable()
	DeleteAllItems()
}

func TestSaveItem(t *testing.T) {
	SaveItem(i)
	nRows := len(GetAllItems())
	if nRows != 1 {
		t.Errorf("Got incorrect number of rows back! Got %d, wanted 1", nRows)
	}
	DeleteAllItems()
}

func TestUpdateItem(t *testing.T) {
	SaveItem(i)
	items := GetAllItems()
	newItem := item.Item{
		Name:   "NEW NAME",
		Note:   "NEW NOTE",
		Amount: 2,
		Status: "new",
		Cat_id: 0,
	}

	UpdateItem(newItem, items[0].ID)

	nRows := len(GetAllItems())
	if nRows != 1 {
		t.Errorf("Got incorrect number of rows back! Inserted new item instead of updating existing one! Got %d, wanted 1", nRows)
	}

	items = GetAllItems()

	if items[0].Name != "NEW NAME" {
		t.Errorf("Got wrong NAME! Wanted %s got %s", newItem.Name, items[0].Name)
	} else if items[0].Note != "NEW NOTE" {
		t.Errorf("Got wrong NOTE! Wanted %s got %s", newItem.Note, items[0].Note)
	} else if items[0].Amount != 2 {
		t.Errorf("Got wrong NOTE! Wanted %d got %d", newItem.Amount, items[0].Amount)
	}
	DeleteAllItems()
}

func TestSwitchItemStatus(t *testing.T) {
	SaveItem(i)
	items := GetAllItems()
	SwitchItemStatus(items[0].ID)

	items = GetAllItems()
	if items[0].Status != "old" {
		t.Errorf("Got wrong status! Wanted %s, got %s", "old", items[0].Status)
	}
	DeleteAllItems()
}

func TestDeleteItem(t *testing.T) {
	SaveItem(i)
	items := GetAllItems()
	DeleteItem(items[0].ID)

	items = GetAllItems()
	if len(GetAllItems()) > 0 {
		t.Errorf("Got wrong amount of remaining items! Wanted 0, got %d", len(items))
	}
}

func TestGetItemsWithStatus(t *testing.T) {
	SaveItem(i)
	items := GetAllItems()
	SwitchItemStatus(items[0].ID)

	items = GetItemsWithStatus("old")
	if len(items) != 1 {
		t.Errorf("Got wrong number of rows! Wanted 1, got %d", len(items))
	}
	DeleteAllItems()
}
