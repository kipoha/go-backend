package store_test

import (
	"testing"

	"github.com/kipoha/go-backend/internal/app/model"
	"github.com/kipoha/go-backend/internal/app/store"
	"github.com/stretchr/testify/assert"
)

var (
	databaseURL = "postgres://postgres:1234@localhost:5432/golang_db?sslmode=disable"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Username:    "kipoha",
		DisplayName: "Kip",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	_, err := s.User().FindByUsername("imseba")
	assert.Error(t, err)

	u, err := s.User().FindByUsername("kipoha")
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

