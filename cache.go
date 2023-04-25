package cache

type Cache struct {
	data map[string]interface{}
}

func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) interface{} {
	exist := c.isExists(key)
	if exist {
		return c.data[key]
	}
	return nil
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

func (c *Cache) isExists(key string) bool {
	_, ok := c.data[key]
	if !ok {
		return false
	}
	return true
}
