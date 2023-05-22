# einkaufsliste-go-gin  - REST package

A REST api for the einkaufsliste-go-gin project.

## Routes

```text
/items/:status      get all items with status
/items/all          get all items regardless of status
/item/new           create a new item, uses json body
/item/:id/update    update an item with given id with new body
/item/:id/switch    switch status of item from old to new and other way around
/item/:id/delete    delete item
```

Updated on the 2023.05.22 (yyyy.mm.dd)

## Entities

### Item

```go
type Item struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Note   string `json:"note"`
	Amount int    `json:"amount"`
	Status string `json:"status"`
	Cat_id int    `json:"cat_Id"`
}
```

#### Example JSON body for Item entity

```json
{
  "id": 0, // can be removed for new item
  "name": "ITEM",
  "note": "NOTE",
  "amount": 1,
  "status": "new",
  "cat_Id": 0
}
```

### Category

```go
type Category struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Counter    int    `json:"counter"`
	Color      string `json:"color"`
	Color_name string `json:"color_Name"`
}
```
