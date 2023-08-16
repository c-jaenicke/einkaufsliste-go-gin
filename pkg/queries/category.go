package queries

import (
	"context"
	"fmt"
	"shopping-list/ent"
	"shopping-list/ent/category"
	"shopping-list/pkg/logging"
)

type CategoryStruct struct {
	Id    int    `json:"id"`    // id of category
	Name  string `json:"name"`  // name of category
	Color string `json:"color"` // hex color code of category
}

func (c *CategoryStruct) Create(ctx context.Context, client *ent.Client) error {
	cat, err := client.Category.Create().SetName(c.Name).SetColor(c.Color).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create new category: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("created new category successfully: %v", cat))
	return nil
}

func (c *CategoryStruct) Delete(ctx context.Context, client *ent.Client) error {
	err := client.Category.DeleteOneID(c.Id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete category with id %d: %w", c.Id, err)
	}

	logging.LogInfo(fmt.Sprintf("category with id %d deleted successfully", c.Id))
	return nil
}

func GetCategoryByName(ctx context.Context, client *ent.Client, name string) (*ent.Category, error) {
	cat, err := client.Category.Query().Where(category.Name(name)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find category with name %s", name)
	}

	logging.LogInfo(fmt.Sprintf("found category with name %s successfully: %v", name, cat))
	return cat, nil
}

func GetCategoryById(ctx context.Context, client *ent.Client, id int) (*ent.Category, error) {
	cat, err := client.Category.Query().Where(category.ID(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find category with id %d: %w", id, err)
	}

	logging.LogInfo(fmt.Sprintf("found category with id %d successfully: %v", id, cat))
	return cat, nil
}

func GetAllCategories(ctx context.Context, client *ent.Client) ([]*ent.Category, error) {
	cats, err := client.Category.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all categories: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("got total of %d categories successfully", len(cats)))
	return cats, nil
}
