package service

import (
	"github.com/breadwithmeth/naliv_go/internal/models"
	"github.com/breadwithmeth/naliv_go/internal/repository"
)

type ItemService struct {
	repo *repository.ItemRepository
}

func NewItemService(repo *repository.ItemRepository) *ItemService {
	if repo == nil {
		panic("item repository is nil")
	}
	return &ItemService{
		repo: repo,
	}
}



func (s *ItemService) GetItems(business_id int) ([]*models.Item, error) {
	
	items, err := s.repo.GetItems(business_id)
	if err != nil {
		return nil, err
	}
	return items, nil
}
