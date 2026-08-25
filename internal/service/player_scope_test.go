package service

import (
	"testing"
	"time"

	"cardgame/internal/sessionpool"
)

func TestPooledPlayerScopeDoesNotCrossRequests(t *testing.T) {
	pool := sessionpool.New()
	svc := NewPlayerScopeService(pool)
	logged := make(chan string, 2)
	releaseAlice := make(chan struct{})
	releaseBob := make(chan struct{})
	svc.Handle("alice", releaseAlice, logged)
	svc.Handle("bob", releaseBob, logged)
	close(releaseAlice)
	select {
	case got := <-logged:
		if got != "alice" {
			t.Errorf("alice 请求的异步日志玩家=%q", got)
		}
	case <-time.After(time.Second):
		t.Fatal("alice 请求的异步日志没有写出")
	}
	scope := pool.Get()
	if scope.PlayerID != "" || len(scope.Labels) != 0 {
		t.Errorf("归还后的请求状态未清空: player=%q labels=%v", scope.PlayerID, scope.Labels)
	}
}
