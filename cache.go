package cache

import (
	"sync"
	"time"
)

type value struct {
	value interface{}
	ttl   *time.Time
}

type Cache struct {
	ticker *time.Ticker
	data   sync.Map
	ttl    time.Duration
}

func New(ttl time.Duration) *Cache {
	return &Cache{
		ticker: time.NewTicker(ttl),
		data:   sync.Map{},
		ttl:    ttl,
	}
}

func (c *Cache) backgroundCache() {
	for {
		<-c.ticker.C
		c.data.Range(func(key, v interface{}) bool {
			value, ok := v.(*value)
			if !ok {
				return true
			}
			if value.ttl == nil {
				return true
			}
			if time.Now().After(*value.ttl) {
				c.data.Delete(key)
			}
			return true
		})
	}
}
func (db *Cache) Get(key interface{}) (result interface{}, ok bool) {
	load, ok := db.data.Load(key)
	if !ok {
		return nil, false
	}

	val, ok := load.(*value)
	if !ok {
		return nil, false
	}

	return val.value, true
}

func (c *Cache) Set(key string, v interface{}) {
	t := time.Now().Add(c.ttl)
	c.data.Store(key, &value{v, &t})
}

func (c *Cache) Delete(key string) {
	c.data.Delete(key)
}
