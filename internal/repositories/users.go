package repositories

import (
	"database/sql"

	"github.com/badimalex/goshop/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
		user.Username,
		user.Password,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(
		"SELECT id, username, password FROM users WHERE username = $1",
		username,
	).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
