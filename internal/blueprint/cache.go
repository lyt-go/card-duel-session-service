package blueprint

import "sync"

type Rules struct{ Allowed map[string]bool }
type Blueprint struct {
	ID    string
	Cards []string
	Rules *Rules
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]*Blueprint
}

func NewCache() *Cache                       { return &Cache{items: make(map[string]*Blueprint)} }
func (c *Cache) Put(id string, b *Blueprint) { c.mu.Lock(); defer c.mu.Unlock(); c.items[id] = b }
func (c *Cache) Get(id string) (*Blueprint, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	b, ok := c.items[id]
	return b, ok
}
