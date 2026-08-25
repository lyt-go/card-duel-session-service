package service

import "cardgame/internal/standings"

type StandingsService struct {
	table *standings.Table
	cache *standings.Cache
}

func NewStandingsService(table *standings.Table, cache *standings.Cache) *StandingsService {
	return &StandingsService{table: table, cache: cache}
}
func (s *StandingsService) Refresh()                      { s.cache.Put(s.table.Snapshot()) }
func (s *StandingsService) CachedScore(player string) int { return s.cache.Score(player) }
