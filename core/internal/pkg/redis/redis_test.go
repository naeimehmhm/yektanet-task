package redis

import (
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedis_Ping(t *testing.T) {
	client, mock := redismock.NewClientMock()
	mock.ExpectPing().SetVal("PONG")
	r := Redis{
		client: client,
	}

	assert.NoError(t, r.Ping())
}