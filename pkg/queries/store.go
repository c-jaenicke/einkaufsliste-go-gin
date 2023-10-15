package queries

import (
	"context"
	"fmt"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent"
	"github.com/c-jaenicke/einkaufsliste-go-gin/ent/store"
	"github.com/c-jaenicke/einkaufsliste-go-gin/pkg/logging"
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
		return fmt.Errorf("failed to delete store with id %d: %w", s.Id, err)
	}

	logging.LogInfo(fmt.Sprintf("store with id %d deleted successfully", s.Id))
	return nil
}

func GetStoreByName(ctx context.Context, client *ent.Client, name string) (*ent.Store, error) {
	st, err := client.Store.Query().Where(store.Name(name)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get store with name %s: %w", name, err)
	}

	logging.LogInfo(fmt.Sprintf("got store with name %s successfully: %v", name, st))
	return st, nil
}

func GetStoreById(ctx context.Context, client *ent.Client, id int) (*ent.Store, error) {
	store, err := client.Store.Query().Where(store.ID(id)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get store with id %d", id)
	}

	logging.LogInfo(fmt.Sprintf("got store with id %d successfully: %v", id, store))
	return store, nil
}

func GetAllStores(ctx context.Context, client *ent.Client) ([]*ent.Store, error) {
	sts, err := client.Store.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all stores: %w", err)
	}

	logging.LogInfo(fmt.Sprintf("got total of %d stores successfully", len(sts)))
	return sts, nil
}
