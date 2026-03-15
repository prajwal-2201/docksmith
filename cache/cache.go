package cache

type Cache struct {
	store map[string]string
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]string),
	}
}

func (c *Cache) Lookup(key string) (string, bool) {
	digest, ok := c.store[key]
	return digest, ok
}

func (c *Cache) Store(key string, digest string) {
	c.store[key] = digest
}
