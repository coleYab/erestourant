package store

import (
	"context"

	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/dto"
	"github.com/google/uuid"
)

type MenuStore struct {
	qry *repository.Queries
}

func NewMenuStore(qry *repository.Queries) *MenuStore {
	return &MenuStore{qry: qry}
}

func (s *MenuStore) CreateMenu(menuParams dto.CreateMenuDto) (repository.MenuItem, error) {
	ctx := context.Background()
	return s.qry.CreateMenu(ctx, repository.CreateMenuParams{
		Name: menuParams.Name,
		Description: menuParams.Description,
		Price: menuParams.Price,
		Qty: menuParams.Qty,
	})
}

func (s *MenuStore) GetMenuByID(id string) (repository.MenuItem, error) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.GetMenuById(ctx, uid)
}


func (s *MenuStore) GetAllMenu() ([]repository.MenuItem, error) {
	ctx := context.Background()
	return s.qry.GetAllMenu(ctx)
}

func (s *MenuStore) DeleteMenu(id string) error {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()
	return s.qry.DeleteMenuById(ctx, uid)
}


