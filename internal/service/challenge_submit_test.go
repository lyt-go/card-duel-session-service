package service

import (
	"errors"
	"testing"

	"cardgame/internal/challengerepo"
	"cardgame/internal/gateway"
)

func TestRejectedChallengeIsNotRetriedOrCommitted(t *testing.T) {
	client := &gateway.Client{}
	repo := &challengerepo.Repository{}
	err := SubmitChallenge("reject", client, repo)
	var remote *gateway.RemoteError
	if !errors.As(err, &remote) || remote.Kind != "rejected" {
		t.Errorf("拒绝错误类型丢失: %v", err)
	}
	if client.Calls != 1 {
		t.Errorf("被拒绝的牌组请求次数=%d", client.Calls)
	}
	if repo.Writes != 0 || repo.Rollbacks != 1 {
		t.Errorf("写入=%d 回滚=%d", repo.Writes, repo.Rollbacks)
	}
}
