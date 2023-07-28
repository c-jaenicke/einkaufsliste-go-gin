package queries

import (
	"context"
	"fmt"
	"shopping-list/ent"
	"shopping-list/ent/item"
	"shopping-list/ent/store"
	"shopping-list/pkg/logging"
)

// TODO add more fields here, store
type ItemStruct struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Note    string `json:"note"`
	Amount  int    `json:"amount"`
	Status  string `json:"status"`
	StoreId int    `json:"store_id"`
}

func (itemStruct *ItemStruct) CreateItem(ctx context.Context, client *ent.Client) error {
	item, err := client.Item.
		Create().SetName(itemStruct.Name).SetNote(itemStruct.Note).SetAmount(itemStruct.Amount).SetStoreID(itemStruct.StoreId).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create a new item: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("created new item: %v", item))
	return nil
}

func (itemStruct *ItemStruct) UpdateItemById(ctx context.Context, client *ent.Client) error {
	item, err := client.Item.UpdateOneID(itemStruct.Id).SetName(itemStruct.Name).SetNote(itemStruct.Note).SetAmount(itemStruct.Amount).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update the item: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("updated the item: %v", item))
	return nil
}

func (itemStruct *ItemStruct) UpdateItemStatus(ctx context.Context, client *ent.Client) error {
	var status string
	if itemStruct.Status == "new" {
		status = "bought"
	} else {
		status = "new"
	}

	item, err := client.Item.UpdateOneID(itemStruct.Id).SetStatus(status).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update Status of item item: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("updated item Status: %v", item))
	return nil
}

func (itemStruct *ItemStruct) DeleteItemById(ctx context.Context, client *ent.Client) error {
	err := client.Item.DeleteOneID(itemStruct.Id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete the item: %w", err)
	}

	logging.LogInfo("item deleted")
	return nil
}

func GetItemById(ctx context.Context, client *ent.Client, id int) (*ent.Item, error) {
	item, err := client.Item.Query().Where(item.ID(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find item : %w", err)
	}

	logging.LogInfo(fmt.Sprintf("found item: %w", item))
	return item, nil
}

func GetAllItems(ctx context.Context, client *ent.Client) ([]*ent.Item, error) {
	items, err := client.Item.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("got items: %d", len(items)))
	return items, nil
}

func GetAllItemsByStatus(ctx context.Context, client *ent.Client, status string) ([]*ent.Item, error) {
	items, err := client.Item.Query().Where(item.Status(status)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get items with Status new: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("found items with Status new: %d", len(items)))
	return items, nil
}

// GetAllItemsByStoreId returns all items that belong to a store
func GetAllItemsByStoreId(ctx context.Context, client *ent.Client, storeId int) ([]*ent.Item, error) {
	items, err := client.Item.Query().Where(item.HasStoreWith(store.ID(storeId))).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get items registered under store %d: %w", storeId, err)
	}

	logging.LogInfo(fmt.Sprintf("found items registered under store %d: %d", storeId, len(items)))
	return items, nil
}

// GetAllItemsByCategoryId returns all items that belong to a category
func GetAllItemsByCategoryId(ctx context.Context, client *ent.Client, categoryId int) ([]*ent.Item, error) {
	items, err := client.Item.Query().Where(item.HasStoreWith(store.ID(categoryId))).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get items registered under store %d: %w", categoryId, err)
	}

	logging.LogInfo(fmt.Sprintf("found items registered under store %d: %d", categoryId, len(items)))
	return items, nil
}
