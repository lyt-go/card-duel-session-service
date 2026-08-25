package service

import (
	"fmt"

	"cardgame/internal/matchstate"
)

type MatchRetry struct {
	store    *matchstate.Store
	effect   *matchstate.Effect
	lateDone chan struct{}
}

func NewMatchRetry(store *matchstate.Store, effect *matchstate.Effect) *MatchRetry {
	return &MatchRetry{store: store, effect: effect, lateDone: make(chan struct{})}
}

func (r *MatchRetry) Run(id string, releaseLate <-chan struct{}) error {
	r.store.WriteDetail(id, matchstate.State{Status: "running", Version: 1})
	go func() {
		defer close(r.lateDone)
		<-releaseLate
		// 首轮回执的延迟回调：用版本守卫写入，绝不能把已到达
		// succeeded 终态的列表覆盖回 running。
		r.store.WriteCacheIfVersion(id, matchstate.State{Status: "running", Version: 1})
	}()
	// 同一对局的动作只能发生一次：重试复用同一对局键，借助 Effect
	// 的去重表保证幂等，首轮回执失败与重试成功只记录一次动作。
	for attempt := 1; attempt <= 2; attempt++ {
		if err := r.effect.Apply(id); err != nil {
			continue
		}
		state := matchstate.State{Status: "succeeded", Version: attempt + 1}
		r.store.WriteDetail(id, state)
		r.store.WriteCache(id, state)
		return nil
	}
	return fmt.Errorf("重试失败")
}
func (r *MatchRetry) WaitLate() { <-r.lateDone }
