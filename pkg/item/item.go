package item

import "fmt"

type Item struct {
	ID     int
	Name   string
	Note   string
	Amount int
	Status string
}

func (item Item) List() string {
	return fmt.Sprintf("%s, %s, %d, %s", item.Name, item.Note, item.Amount, item.Status)
}

func (item Item) DebugList() string {
	return fmt.Sprintf("(%d) %s, %s, %d, %s", item.ID, item.Name, item.Note, item.Amount, item.Status)
}
