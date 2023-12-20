package model

type ICacher interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type MemCache struct {
	data map[string]interface{}
}

func (c *MemCache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *MemCache) Get(key string) (interface{}, bool) {
	if c.data == nil {
		return nil, false
	}

	value, ok := c.data[key]
	return value, ok
}
