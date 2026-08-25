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

// Snapshot 返回当前分数的拷贝。调用方拿到的是独立副本，后续对 Table 的更新
// 不会波及快照，快照本身也不会与 Table 共享底层 map，从而避免并发读写竞态。
func (t *Table) Snapshot() map[string]int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	snapshot := make(map[string]int, len(t.scores))
	for k, v := range t.scores {
		snapshot[k] = v
	}
	return snapshot
}

// Cache 保存某次 Refresh 时的分数快照。快照是冻结的：Table 之后的更新不会
// 改变已缓存的副本，只有在主动调用 Put（刷新）时才会替换为新的拷贝。
type Cache struct {
	mu     sync.RWMutex
	scores map[string]int
}

// Put 用 scores 的拷贝替换缓存。拷贝一份是为了让缓存持有一个不被外部调用方
// 持有的独立快照，确保快照内容在下次刷新前保持不变。
func (c *Cache) Put(scores map[string]int) {
	snapshot := make(map[string]int, len(scores))
	for k, v := range scores {
		snapshot[k] = v
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.scores = snapshot
}

func (c *Cache) Score(player string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.scores[player]
}
