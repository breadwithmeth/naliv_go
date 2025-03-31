package repository

import (
	"database/sql"

	"github.com/breadwithmeth/naliv_go/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	if db == nil {
		panic("database connection is nil")
	}
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	rows, err := r.db.Query("SELECT user_id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
