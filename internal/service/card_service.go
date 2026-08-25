package service

import (
	"sort"
	"time"

	"cardgame/internal/model"
	"cardgame/pkg/idgen"
)

func (s *Service) CreateCard(input model.Card) (*model.Card, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetCardByName(input.Name); err == nil {
		return nil, model.NewValidationError("name", "已存在")
	}
	now := time.Now()
	input.ID = idgen.Hex()
	input.CreatedAt = now
	input.UpdatedAt = now
	if err := s.store.CreateCard(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetCard(id string) (*model.Card, error) {
	return s.store.GetCard(id)
}

func (s *Service) ListCards(filter model.CardFilter, page, size int) ([]*model.Card, int, error) {
	all := s.store.ListCards()
	matched := make([]*model.Card, 0, len(all))
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
		return []*model.Card{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) UpdateCard(id string, input model.Card) (*model.Card, error) {
	exist, err := s.store.GetCard(id)
	if err != nil {
		return nil, err
	}
	exist.Name = input.Name
	exist.Description = input.Description
	exist.Cost = input.Cost
	exist.Attack = input.Attack
	exist.Health = input.Health
	exist.Rarity = input.Rarity
	exist.Type = input.Type
	if err := exist.Validate(); err != nil {
		return nil, err
	}
	exist.UpdatedAt = time.Now()
	if err := s.store.UpdateCard(exist); err != nil {
		return nil, err
	}
	return exist, nil
}

func (s *Service) DeleteCard(id string) error {
	return s.store.DeleteCard(id)
}
