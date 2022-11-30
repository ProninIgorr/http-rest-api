package sqlstore_test

import (
	"github.com/ProninIgorr/http-rest-api/internal/app/model"
	"github.com/ProninIgorr/http-rest-api/internal/app/store"
	"github.com/ProninIgorr/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURl)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURl)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@email.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURl)
	defer teardown("users")

	s := sqlstore.New(db)
	id := 21
	_, err := s.User().Find(id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.ID = id
	s.User().Create(u)
	u, err = s.User().Find(id)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
