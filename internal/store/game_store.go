package store

import (
	"cardgame/internal/model"
)

func (s *MemoryStore) CreateGame(x *model.Game) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.games[x.ID] = x
	return nil
}

func (s *MemoryStore) GetGame(id string) (*model.Game, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	x, ok := s.games[id]
	if !ok {
		return nil, ErrNotFound
	}
	return x, nil
}

func (s *MemoryStore) ListGames() []*model.Game {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Game, 0, len(s.games))
	for _, x := range s.games {
		list = append(list, x)
	}
	return list
}

func (s *MemoryStore) UpdateGame(x *model.Game) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.games[x.ID]; !ok {
		return ErrNotFound
	}
	s.games[x.ID] = x
	return nil
}

func (s *MemoryStore) DeleteGame(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.games[id]; !ok {
		return ErrNotFound
	}
	delete(s.games, id)
	return nil
}
