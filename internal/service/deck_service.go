package service

import (
	"sort"
	"time"

	"cardgame/internal/model"
	"cardgame/pkg/idgen"
)

func (s *Service) CreateDeck(input model.Deck) (*model.Deck, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetDeckByName(input.Name); err == nil {
		return nil, model.NewValidationError("name", "已存在")
	}
	if input.PlayerID != "" {
		if _, err := s.store.GetPlayer(input.PlayerID); err != nil {
			return nil, model.NewValidationError("player_id", "关联的玩家不存在")
		}
	}
	now := time.Now()
	input.ID = idgen.Hex()
	input.CreatedAt = now
	input.UpdatedAt = now
	if err := s.store.CreateDeck(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetDeck(id string) (*model.Deck, error) {
	return s.store.GetDeck(id)
}

func (s *Service) ListDecks(filter model.DeckFilter, page, size int) ([]*model.Deck, int, error) {
	all := s.store.ListDecks()
	matched := make([]*model.Deck, 0, len(all))
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
		return []*model.Deck{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateDeck(id string, input model.Deck) (*model.Deck, error) {
	exist, err := s.store.GetDeck(id)
	if err != nil {
		return nil, err
	}
	exist.Name = input.Name
	exist.Description = input.Description
	exist.CardCount = input.CardCount
	exist.PlayerID = input.PlayerID
	if err := exist.Validate(); err != nil {
		return nil, err
	}
	exist.UpdatedAt = time.Now()
	if err := s.store.UpdateDeck(exist); err != nil {
		return nil, err
	}
	return exist, nil
}

func (s *Service) DeleteDeck(id string) error {
	return s.store.DeleteDeck(id)
}

func (s *Service) TransitionDeck(id, target string) (*model.Deck, error) {
	if !model.DeckValidStatus(target) {
		return nil, model.NewValidationError("status", "目标状态不合法")
	}
	exist, err := s.store.GetDeck(id)
	if err != nil {
		return nil, err
	}
	if !model.DeckCanTransition(exist.Status, target) {
		return nil, model.NewValidationError("status", "不允许从 "+exist.Status+" 流转到 "+target)
	}
	exist.Status = target
	exist.UpdatedAt = time.Now()
	if err := s.store.UpdateDeck(exist); err != nil {
		return nil, err
	}
	return exist, nil
}
