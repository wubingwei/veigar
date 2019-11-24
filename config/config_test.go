package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	assert.NotNil(t, VeigarRoot)

	assert.NotNil(t, DevMode)

	assert.NotNil(t, Mysql.Host)
	assert.NotNil(t, Mysql.Port)
	assert.NotNil(t, Mysql.User)
	assert.NotNil(t, Mysql.Password)
	assert.NotNil(t, Mysql.DataBase)
}
