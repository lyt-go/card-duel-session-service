package service

import (
	"sort"
	"time"

	"cardgame/internal/model"
	"cardgame/pkg/idgen"
)

func (s *Service) CreateCollection(input model.Collection) (*model.Collection, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetPlayer(input.PlayerID); err != nil {
		return nil, model.NewValidationError("player_id", "关联的玩家不存在")
	}
	if _, err := s.store.GetCard(input.CardID); err != nil {
		return nil, model.NewValidationError("card_id", "关联的卡牌不存在")
	}
	now := time.Now()
	input.ID = idgen.Hex()
	input.CreatedAt = now
	if err := s.store.CreateCollection(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetCollection(id string) (*model.Collection, error) {
	return s.store.GetCollection(id)
}

func (s *Service) ListCollections(filter model.CollectionFilter, page, size int) ([]*model.Collection, int, error) {
	all := s.store.ListCollections()
	matched := make([]*model.Collection, 0, len(all))
	for _, x := range all {
		if filter.Match(x) {
			matched = append(matched, x)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].CreatedAt.After(matched[j].CreatedAt)
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.Collection{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) DeleteCollection(id string) error {
	return s.store.DeleteCollection(id)
}
