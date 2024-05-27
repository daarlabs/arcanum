package auth

import (
	"context"
	"testing"
	
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	
	"github.com/daarlabs/arcanum/quirk"
)

func createTestDatabaseConnection(t *testing.T) *quirk.DB {
	db, err := quirk.Connect(
		quirk.WithPostgres(),
		quirk.WithHost("localhost"),
		quirk.WithPort(5432),
		quirk.WithDbname("test"),
		quirk.WithUser("cream"),
		quirk.WithPassword("cream"),
		quirk.WithSslDisable(),
	)
	assert.NoError(t, err)
	assert.NoError(t, db.Ping())
	return db
}

func createTestRedisConnection(t *testing.T) *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
			DB:   10,
		},
	)
	assert.Nil(t, client.Ping(context.Background()).Err())
	return client
}
