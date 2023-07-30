package queries

import (
	"context"
	"fmt"
	"shopping-list/ent"
	"shopping-list/ent/store"
	"shopping-list/pkg/logging"
)

type StoreStruct struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (s *StoreStruct) Create(ctx context.Context, client *ent.Client) error {
	st, err := client.Store.Create().SetName(s.Name).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create new store: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("created new store successfully: %v", st))
	return nil
}

func (s *StoreStruct) Delete(ctx context.Context, client *ent.Client) error {
	err := client.Store.DeleteOneID(s.Id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete store id: %w", err)
	}

	logging.LogInfo("store deleted successfully")
	return nil
}

func GetStoreByName(ctx context.Context, client *ent.Client, name string) (*ent.Store, error) {
	st, err := client.Store.Query().Where(store.Name(name)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find store with name %s", name)
	}

	logging.LogInfo(fmt.Sprintf("found store successfully: %v", st))
	return st, nil
}

func GetStoreById(ctx context.Context, client *ent.Client, id int) (*ent.Store, error) {
	store, err := client.Store.Query().Where(store.ID(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find store with id %d", id)
	}

	logging.LogInfo(fmt.Sprintf("found store successfully: %v", store))
	return store, nil
}

func GetAllStores(ctx context.Context, client *ent.Client) ([]*ent.Store, error) {
	sts, err := client.Store.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all stores: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("got stores: %d", len(sts)))
	return sts, nil
}
