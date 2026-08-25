package store

import (
	"cardgame/internal/model"
)

func (s *MemoryStore) CreateCollection(x *model.Collection) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.collections[x.ID] = x
	return nil
}

func (s *MemoryStore) GetCollection(id string) (*model.Collection, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	x, ok := s.collections[id]
	if !ok {
		return nil, ErrNotFound
	}
	return x, nil
}

func (s *MemoryStore) ListCollections() []*model.Collection {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.Collection, 0, len(s.collections))
	for _, x := range s.collections {
		list = append(list, x)
	}
	return list
}

func (s *MemoryStore) DeleteCollection(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.collections[id]; !ok {
		return ErrNotFound
	}
	delete(s.collections, id)
	return nil
}
