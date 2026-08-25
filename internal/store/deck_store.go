package store

import (
	"cardgame/internal/model"
)

func (s *MemoryStore) CreateDeck(x *model.Deck) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.decks {
		if exist.Name == x.Name {
			return ErrConflict
		}
	}
	s.decks[x.ID] = x
	return nil
}

func (s *MemoryStore) GetDeck(id string) (*model.Deck, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	x, ok := s.decks[id]
	if !ok {
		return nil, ErrNotFound
	}
	return x, nil
}

func (s *MemoryStore) GetDeckByName(v string) (*model.Deck, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, x := range s.decks {
		if x.Name == v {
			return x, nil
		}
	}
	return nil, ErrNotFound
}

func (s *MemoryStore) ListDecks() []*model.Deck {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Deck, 0, len(s.decks))
	for _, x := range s.decks {
		list = append(list, x)
	}
	return list
}

func (s *MemoryStore) UpdateDeck(x *model.Deck) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.decks[x.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.decks {
		if exist.ID != x.ID && exist.Name == x.Name {
			return ErrConflict
		}
	}
	s.decks[x.ID] = x
	return nil
}

func (s *MemoryStore) DeleteDeck(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.decks[id]; !ok {
		return ErrNotFound
	}
	delete(s.decks, id)
	return nil
}
