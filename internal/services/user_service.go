package service

import (
	"errors"

	"github.com/breadwithmeth/naliv_go/internal/models"
	"github.com/breadwithmeth/naliv_go/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	if repo == nil {
		panic("user repository is nil")
	}

	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	// users := []*models.User{}
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) ValidateToken(token string) (int, error) {
	userID, err := s.repo.ValidateToken(token)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (s *UserService) getUserByToken(token string) (*models.User, error) {
	userID, err := s.ValidateToken(token)
	if err != nil {
		return nil, err
	}
	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserAddresses(userID int) ([]*models.Address, error) {
	addresses, err := s.repo.GetUserAddresses(userID)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (s *UserService) GetUserByToken(token string) (*models.User, error) {
	user, err := s.getUserByToken(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}
