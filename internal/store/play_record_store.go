package store

import (
	"cardgame/internal/model"
)

func (s *MemoryStore) CreatePlayRecord(x *model.PlayRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.playRecords[x.ID] = x
	return nil
}

func (s *MemoryStore) GetPlayRecord(id string) (*model.PlayRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	x, ok := s.playRecords[id]
	if !ok {
		return nil, ErrNotFound
	}
	return x, nil
}

func (s *MemoryStore) ListPlayRecords() []*model.PlayRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*model.PlayRecord, 0, len(s.playRecords))
	for _, x := range s.playRecords {
		list = append(list, x)
	}
	return list
}

func (s *MemoryStore) DeletePlayRecord(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.playRecords[id]; !ok {
		return ErrNotFound
	}
	delete(s.playRecords, id)
	return nil
}
