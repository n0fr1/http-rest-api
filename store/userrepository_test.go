package store_test

import (
	"testing"

	"github.com/n0fr1/http-rest-api/model"
	"github.com/n0fr1/http-rest-api/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {

	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
