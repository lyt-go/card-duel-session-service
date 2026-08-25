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
		r.store.WriteCache(id, matchstate.State{Status: "running", Version: 1})
	}()
	for attempt := 1; attempt <= 2; attempt++ {
		if err := r.effect.Apply(fmt.Sprintf("%s-%d", id, attempt)); err != nil {
			continue
		}
		state := matchstate.State{Status: "succeeded", Version: attempt}
		r.store.WriteDetail(id, state)
		r.store.WriteCache(id, state)
		return nil
	}
	return fmt.Errorf("重试失败")
}
func (r *MatchRetry) WaitLate() { <-r.lateDone }
