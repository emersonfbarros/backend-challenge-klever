package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	cache := &MemCache{data: make(map[string]interface{})}
	cache.Set("key", "value")

	value, ok := cache.data["key"]
	assert.True(t, ok)
	assert.Equal(t, "value", value)
}

func TestGet(t *testing.T) {
	cache := &MemCache{data: make(map[string]interface{})}
	cache.data["key"] = "value"

	value, ok := cache.Get("key")
	assert.True(t, ok)
	assert.Equal(t, "value", value)
}

func TestGetNoKey(t *testing.T) {
	cache := &MemCache{data: make(map[string]interface{})}

	_, ok := cache.Get("key")
	assert.False(t, ok)
}
