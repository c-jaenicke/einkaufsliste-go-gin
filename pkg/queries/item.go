package queries

import (
	"context"
	"fmt"
	"shopping-list/ent"
	"shopping-list/ent/category"
	"shopping-list/ent/item"
	"shopping-list/ent/store"
	"shopping-list/pkg/logging"
)

type ItemStruct struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Note       string `json:"note"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	StoreId    int    `json:"store_id"`
	CategoryId int    `json:"category_id"`
}

func (itemStruct *ItemStruct) Create(ctx context.Context, client *ent.Client) error {
	it, err := client.Item.
		Create().
		SetName(itemStruct.Name).
		SetNote(itemStruct.Note).
		SetAmount(itemStruct.Amount).
		SetStoreID(itemStruct.StoreId).
		SetCategoryID(itemStruct.CategoryId).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create a new item: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("created new item: %v", it))
	return nil
}

func (itemStruct *ItemStruct) Update(ctx context.Context, client *ent.Client) error {
	it, err := client.Item.
		UpdateOneID(itemStruct.Id).
		SetName(itemStruct.Name).
		SetNote(itemStruct.Note).
		SetAmount(itemStruct.Amount).
		SetStoreID(itemStruct.StoreId).
		SetCategoryID(itemStruct.CategoryId).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update the item: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("updated the item: %v", it))
	return nil
}

// SwitchItemStatus switches status of the item with the given id, from "new" to "bought" and "bought" to "new"
func SwitchItemStatus(ctx context.Context, client *ent.Client, id int) error {
	it, err := GetItemById(context.Background(), client, id)
	if err != nil {
		return fmt.Errorf("item not found: %w", err)
	}
	var status string
	if it.Status == "new" {
		status = "bought"
	} else {
		status = "new"
	}

	it, err = client.Item.UpdateOneID(id).SetStatus(status).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update Status of item item: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("updated item Status: %v", it))
	return nil
}

func (itemStruct *ItemStruct) Delete(ctx context.Context, client *ent.Client) error {
	err := client.Item.DeleteOneID(itemStruct.Id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete the item: %w", err)
	}

	logging.LogInfo("item deleted")
	return nil
}

func GetItemById(ctx context.Context, client *ent.Client, id int) (*ent.Item, error) {
	it, err := client.Item.Query().Where(item.ID(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find item : %w", err)
	}

	logging.LogInfo(fmt.Sprintf("found item: %w", it))
	return it, nil
}

func GetAllItems(ctx context.Context, client *ent.Client) ([]*ent.Item, error) {
	its, err := client.Item.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("got items: %d", len(its)))
	return its, nil
}

func GetAllItemsByStatus(ctx context.Context, client *ent.Client, status string) ([]*ent.Item, error) {
	its, err := client.Item.Query().Where(item.Status(status)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get items with Status new: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("found items with Status new: %d", len(its)))
	return its, nil
}

// GetAllItemsByStoreId returns all items that belong to a store
func GetAllItemsByStoreId(ctx context.Context, client *ent.Client, storeId int) ([]*ent.Item, error) {
	its, err := client.Item.Query().Where(item.HasStoreWith(store.ID(storeId))).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get items registered under store %d: %w", storeId, err)
	}

	logging.LogInfo(fmt.Sprintf("found items registered under store %d: %d", storeId, len(its)))
	return its, nil
}

// GetAllItemsByCategoryId returns all items that belong to a category
func GetAllItemsByCategoryId(ctx context.Context, client *ent.Client, categoryId int) ([]*ent.Item, error) {
	its, err := client.Item.Query().Where(item.HasCategoryWith(category.ID(categoryId))).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get items registered under store %d: %w", categoryId, err)
	}

	logging.LogInfo(fmt.Sprintf("found items registered under store %d: %d", categoryId, len(its)))
	return its, nil
}
