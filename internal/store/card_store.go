package store

import (
	"cardgame/internal/model"
)

func (s *MemoryStore) CreateCard(x *model.Card) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, exist := range s.cards {
		if exist.Name == x.Name {
			return ErrConflict
		}
	}
	s.cards[x.ID] = x
	return nil
}

func (s *MemoryStore) GetCard(id string) (*model.Card, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	x, ok := s.cards[id]
	if !ok {
		return nil, ErrNotFound
	}
	return x, nil
}

func (s *MemoryStore) GetCardByName(v string) (*model.Card, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, x := range s.cards {
		if x.Name == v {
			return x, nil
		}
	}
	return nil, ErrNotFound
}

func (s *MemoryStore) ListCards() []*model.Card {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Card, 0, len(s.cards))
	for _, x := range s.cards {
		list = append(list, x)
	}
	return list
}

func (s *MemoryStore) UpdateCard(x *model.Card) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.cards[x.ID]; !ok {
		return ErrNotFound
	}
	for _, exist := range s.cards {
		if exist.ID != x.ID && exist.Name == x.Name {
			return ErrConflict
		}
	}
	s.cards[x.ID] = x
	return nil
}

func (s *MemoryStore) DeleteCard(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.cards[id]; !ok {
		return ErrNotFound
	}
	delete(s.cards, id)
	return nil
}
