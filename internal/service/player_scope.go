package service

import "cardgame/internal/sessionpool"

type PlayerScopeService struct{ pool *sessionpool.Pool }

func NewPlayerScopeService(pool *sessionpool.Pool) *PlayerScopeService {
	return &PlayerScopeService{pool: pool}
}

func (s *PlayerScopeService) Handle(player string, release <-chan struct{}, logged chan<- string) {
	scope := s.pool.Get()
	scope.PlayerID = player
	scope.Labels = append(scope.Labels, "match-feed")
	// Capture the player identity by value: the scope is returned to the pool
	// below and may be reused (and its PlayerID overwritten) by the next
	// request before this goroutine finishes its async log. Reading the
	// captured value keeps different players' requests fully isolated.
	playerID := scope.PlayerID
	go func() { <-release; logged <- playerID }()
	s.pool.Put(scope)
}
