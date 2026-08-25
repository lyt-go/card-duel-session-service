package store

import (
	"cardgame/internal/model"
)

func (s *MemoryStore) CreatePlayer(x *model.Player) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.players {
		if exist.Nickname == x.Nickname {
			return ErrConflict
		}
	}
	s.players[x.ID] = x
	return nil
}

func (s *MemoryStore) GetPlayer(id string) (*model.Player, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	x, ok := s.players[id]
	if !ok {
		return nil, ErrNotFound
	}
	return x, nil
}

func (s *MemoryStore) GetPlayerByNickname(v string) (*model.Player, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, x := range s.players {
		if x.Nickname == v {
			return x, nil
		}
	}
	return nil, ErrNotFound
}

func (s *MemoryStore) ListPlayers() []*model.Player {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Player, 0, len(s.players))
	for _, x := range s.players {
		list = append(list, x)
	}
	return list
}

func (s *MemoryStore) UpdatePlayer(x *model.Player) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.players[x.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.players {
		if exist.ID != x.ID && exist.Nickname == x.Nickname {
			return ErrConflict
		}
	}
	s.players[x.ID] = x
	return nil
}

func (s *MemoryStore) DeletePlayer(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.players[id]; !ok {
		return ErrNotFound
	}
	delete(s.players, id)
	return nil
}
