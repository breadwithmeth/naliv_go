package service

import (
	"github.com/breadwithmeth/naliv_go/internal/models"
	"github.com/breadwithmeth/naliv_go/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	if repo == nil {
		panic("item repository is nil")
	}
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) GetCategories(business_id int) ([]*models.Category, error) {
	categories, err := s.repo.GetCategories(business_id)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
