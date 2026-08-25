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
	go func() { <-release; logged <- scope.PlayerID }()
	s.pool.Put(scope)
}
