package cache

import (
	"sync"
)

// Cache represents a thread-safe in-memory cache using sync.Map.
// sync.Map is used to ensure concurrent access to the cache is safe.
type Cache struct {
	store sync.Map
}

// NewCache initializes and returns a new Cache instance.
func NewCache() *Cache {
	return &Cache{}
}

// Get retrieves the value associated with the given key from the cache.
// Returns the value and a boolean indicating whether the key exists in the cache.
func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.store.Load(key) // Retrieve the value from the sync.Map
	if ok {
		// If the key exists, type assert the value to string and return it.
		return value.(string), true
	}
	// If the key does not exist, return an empty string and false.
	return "", false
}

// Set adds a key-value pair to the cache.
// If the key already exists, the value is updated.
func (c *Cache) Set(key, value string) {
	c.store.Store(key, value)
}

// Clear removes all entries from the cache.
// This is done by reinitializing the sync.Map to ensure thread safety.
func (c *Cache) Clear() {
	c.store = sync.Map{} // Replace the current map with a new empty one
}
