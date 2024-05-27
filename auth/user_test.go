package auth

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/daarlabs/arcanum/quirk"
)

func TestUser(t *testing.T) {
	db := createTestDatabaseConnection(t)
	assert.NotNil(t, db)
	assert.NoError(t, DropTable(quirk.New(db)))
	assert.NoError(t, CreateTable(quirk.New(db)))
	um := CreateUserManager(db, nil, 0, "")
	t.Cleanup(
		func() {
			assert.NoError(t, DropTable(quirk.New(db)))
		},
	)
	t.Run(
		"create", func(t *testing.T) {
			id, err := um.Create(
				User{
					Active:   true,
					Roles:    []string{"owner"},
					Email:    "dominik@linduska.dev",
					Password: "123456789",
				},
			)
			assert.NoError(t, err)
			assert.True(t, id > 0)
		},
	)
	t.Run(
		"get", func(t *testing.T) {
			u, err := um.Get()
			assert.NoError(t, err)
			assert.True(t, u.Id > 0)
		},
	)
	t.Run(
		"update", func(t *testing.T) {
			data := User{
				Roles: []string{"admin"},
			}
			assert.NoError(t, um.Update(data, "roles"))
			assert.Equal(t, []string{"admin"}, um.MustGet().Roles)
		},
	)
	t.Run(
		"disable enable", func(t *testing.T) {
			assert.True(t, um.MustGet().Active)
			assert.NoError(t, um.Disable())
			assert.False(t, um.MustGet().Active)
			assert.NoError(t, um.Enable())
			assert.True(t, um.MustGet().Active)
		},
	)
}
