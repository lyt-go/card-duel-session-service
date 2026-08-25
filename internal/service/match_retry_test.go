package service

import (
	"testing"

	"cardgame/internal/matchstate"
)

func TestRetryKeepsSuccessfulMatchState(t *testing.T) {
	store := matchstate.NewStore()
	effect := &matchstate.Effect{}
	retry := NewMatchRetry(store, effect)
	release := make(chan struct{})
	if err := retry.Run("match-7", release); err != nil {
		t.Fatalf("重试返回错误: %v", err)
	}
	close(release)
	retry.WaitLate()
	if len(effect.Actions) != 1 {
		t.Errorf("对局副作用执行了 %d 次: %v", len(effect.Actions), effect.Actions)
	}
	if got := store.Detail("match-7"); got.Status != "succeeded" {
		t.Errorf("详情状态=%s version=%d", got.Status, got.Version)
	}
	if got := store.Listed("match-7"); got.Status != "succeeded" {
		t.Errorf("列表状态=%s version=%d", got.Status, got.Version)
	}
}
