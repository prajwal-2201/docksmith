package cache

type Cache struct{}

func (c *Cache) Lookup(key string) (string, bool) {
	return "", false
}

func (c *Cache) Store(key string, digest string) {}
