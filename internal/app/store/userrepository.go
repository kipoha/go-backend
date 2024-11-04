package store

import "github.com/kipoha/go-backend/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (username, displayName) VALUES ($1, $2) RETURNING id",
		u.Username, u.DisplayName,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, username, displayName FROM users WHERE username = $1",
		username,
	).Scan(&u.ID, &u.Username, &u.DisplayName); err != nil {
		return nil, err
	}
	return u, nil
}
