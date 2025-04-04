package repository

import (
	"database/sql"
	"errors"
	"log"

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

func (r *UserRepository) ValidateToken(token string) (int, error) {
	var userID int
	query := "SELECT user_id FROM users_tokens WHERE token = ?"
	err := r.db.QueryRow(query, token).Scan(&userID)
	log.Println("Validating token:", token)
	log.Println("User ID from token:", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("invalid token")
		}
		return 0, err
	}
	return userID, nil
}
