package queries

import (
	"context"
	"fmt"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent/category"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent/item"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent/store"
	"github.com/c-jaenicke/einkaufsliste-go-gin/pkg/logging"
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

func (item *ItemStruct) Create(ctx context.Context, client *ent.Client) error {
	it, err := client.Item.
		Create().
		SetName(item.Name).
		SetNote(item.Note).
		SetAmount(item.Amount).
		SetStoreID(item.StoreId).
		SetCategoryID(item.CategoryId).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create new item: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("created new item successfully: %v", it))
	return nil
}

func (item *ItemStruct) Update(ctx context.Context, client *ent.Client) error {
	it, err := client.Item.
		UpdateOneID(item.Id).
		SetName(item.Name).
		SetNote(item.Note).
		SetAmount(item.Amount).
		SetStoreID(item.StoreId).
		SetCategoryID(item.CategoryId).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update item with id %d: %w", item.Id, err)
	}

	logging.LogInfo(fmt.Sprintf("updated item with id %d successfully: %v", item.Id, it))
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
		return fmt.Errorf("failed to switch status of item with id %d: %w", id, err)
	}

	logging.LogInfo(fmt.Sprintf("switched status of item with id %d successfully: %v", id, it))
	return nil
}

func (item *ItemStruct) Delete(ctx context.Context, client *ent.Client) error {
	err := client.Item.DeleteOneID(item.Id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete item wit id %d: %w", item.Id, err)
	}

	logging.LogInfo(fmt.Sprintf("item with id %d deleted successfully", item.Id))
	return nil
}

func GetItemById(ctx context.Context, client *ent.Client, id int) (*ent.Item, error) {
	it, err := client.Item.Query().Where(item.ID(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find item with id %d: %w", id, err)
	}

	logging.LogInfo(fmt.Sprintf("found item with id %d successfully: %v", id, it))
	return it, nil
}

func GetAllItems(ctx context.Context, client *ent.Client) ([]*ent.Item, error) {
	its, err := client.Item.Query().Order(item.ByStoreID(), item.ByCategoryID(), item.ByName()).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("got total of %d items successfully", len(its)))
	return its, nil
}

func GetAllItemsByStatus(ctx context.Context, client *ent.Client, status string) ([]*ent.Item, error) {
	its, err := client.Item.Query().Where(item.Status(status)).Order(item.ByStoreID(), item.ByCategoryID(), item.ByName()).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items with status %s: %w", status, err)
	}

	logging.LogInfo(fmt.Sprintf("found total of %d items with status %s successfully", len(its), status))
	return its, nil
}

// GetAllItemsByStoreId returns all items that belong to a store
func GetAllItemsByStoreId(ctx context.Context, client *ent.Client, storeId int) ([]*ent.Item, error) {
	its, err := client.Item.Query().Where(item.HasStoreWith(store.ID(storeId))).Order(item.ByStoreID(), item.ByCategoryID(), item.ByName()).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items registered under store with id %d: %w", storeId, err)
	}

	logging.LogInfo(fmt.Sprintf("got total of %d items registered under store with id %d successfully", storeId, len(its)))
	return its, nil
}

// GetAllItemsByStoreName returns all items that belong to a store
func GetAllItemsByStoreName(ctx context.Context, client *ent.Client, name string) ([]*ent.Item, error) {
	its, err := client.Item.Query().Where(item.HasStoreWith(store.Name(name))).Order(item.ByStoreID(), item.ByCategoryID(), item.ByName()).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items registered under store with name %s: %w", name, err)
	}

	logging.LogInfo(fmt.Sprintf("got total of %d items registered under store with name %s successfully", len(its), name))
	return its, nil
}

// GetAllItemsByCategoryId returns all items that belong to a category
func GetAllItemsByCategoryId(ctx context.Context, client *ent.Client, categoryId int) ([]*ent.Item, error) {
	its, err := client.Item.Query().Where(item.HasCategoryWith(category.ID(categoryId))).Order(item.ByStoreID(), item.ByCategoryID(), item.ByName()).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all items registered under store with id %d: %w", categoryId, err)
	}

	logging.LogInfo(fmt.Sprintf("got total of %d items registered under store with id %d successfully", len(its), categoryId))
	return its, nil
}

// GetAllItemsByCategoryName returns all items that belong to a category
func GetAllItemsByCategoryName(ctx context.Context, client *ent.Client, name string) ([]*ent.Item, error) {
	its, err := client.Item.Query().Where(item.HasCategoryWith(category.Name(name))).Order(item.ByStoreID(), item.ByCategoryID(), item.ByName()).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get items registered under store %s: %w", name, err)
	}

	logging.LogInfo(fmt.Sprintf("got total of %d items registered under store with name %s successfully", len(its), name))
	return its, nil
}
