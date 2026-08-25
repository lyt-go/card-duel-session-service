package standings

import "sync"

type Table struct {
	mu     sync.RWMutex
	scores map[string]int
}

func NewTable() *Table { return &Table{scores: make(map[string]int)} }
func (t *Table) Add(player string, delta int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.scores[player] += delta
}
func (t *Table) Snapshot() map[string]int { t.mu.RLock(); defer t.mu.RUnlock(); return t.scores }

type Cache struct {
	mu     sync.RWMutex
	scores map[string]int
}

func (c *Cache) Put(scores map[string]int) { c.mu.Lock(); defer c.mu.Unlock(); c.scores = scores }
func (c *Cache) Score(player string) int   { c.mu.RLock(); defer c.mu.RUnlock(); return c.scores[player] }
