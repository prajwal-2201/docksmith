package cache

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

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

func ComputeKey(prev string, inst string, workdir string, env map[string]string, files []string) string {

	var parts []string

	// previous layer
	parts = append(parts, prev)

	// instruction text
	parts = append(parts, inst)

	// workdir
	parts = append(parts, workdir)

	// ENV (sorted)
	var envPairs []string
	for k, v := range env {
		envPairs = append(envPairs, fmt.Sprintf("%s=%s", k, v))
	}

	sort.Strings(envPairs)
	parts = append(parts, strings.Join(envPairs, ","))

	// file list (sorted already)
	parts = append(parts, strings.Join(files, ","))

	data := strings.Join(parts, "|")

	hash := sha256.Sum256([]byte(data))

	return fmt.Sprintf("%x", hash)
}
