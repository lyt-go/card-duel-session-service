package service

import (
	"sync"
	"testing"

	"cardgame/internal/standings"
)

func TestStandingsSnapshotStaysStableDuringUpdates(t *testing.T) {
	table := standings.NewTable()
	cache := &standings.Cache{}
	svc := NewStandingsService(table, cache)
	table.Add("alice", 10)
	svc.Refresh()
	table.Add("alice", 5)
	if got := svc.CachedScore("alice"); got != 10 {
		t.Errorf("未刷新时榜单分数=%d, want 10", got)
	}

	snapshot := table.Snapshot()
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); <-start; table.Add("alice", 1) }()
	go func() { defer wg.Done(); <-start; _ = snapshot["alice"] }()
	close(start)
	wg.Wait()
}
