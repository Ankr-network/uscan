package utils

import lru "github.com/hashicorp/golang-lru"

const maxCacheSize = 10240

type Cache struct {
	cache *lru.Cache
}

func NewCache() *Cache {
	c, err := lru.New(maxCacheSize)
	if err != nil {
		panic(err)
	}
	return &Cache{
		cache: c,
	}
}

func (c *Cache) Add(key, value interface{}) bool {
	return c.cache.Add(key, value)
}

func (c *Cache) Remove(key interface{}) bool {
	return c.cache.Remove(key)
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	return c.cache.Get(key)
}
