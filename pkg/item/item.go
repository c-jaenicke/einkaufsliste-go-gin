package item

type Item struct {
	ID     int    // id of an item
	Name   string // main description of an item
	Note   string // note on the item itself
	Amount int    // how much of an item should be bought
	Status string // status of item, can be [new old deleted] can be expanded throughout development
}
