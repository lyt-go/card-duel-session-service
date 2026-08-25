package service

import (
	"sort"
	"time"

	"cardgame/internal/model"
	"cardgame/pkg/idgen"
)

func (s *Service) CreatePlayRecord(input model.PlayRecord) (*model.PlayRecord, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.store.GetGame(input.GameID); err != nil {
		return nil, model.NewValidationError("game_id", "关联的对局不存在")
	}
	if _, err := s.store.GetDeck(input.DeckID); err != nil {
		return nil, model.NewValidationError("deck_id", "关联的牌组不存在")
	}
	if _, err := s.store.GetCard(input.CardID); err != nil {
		return nil, model.NewValidationError("card_id", "关联的卡牌不存在")
	}
	now := time.Now()
	input.ID = idgen.Hex()
	input.CreatedAt = now
	if input.PlayedAt == 0 {
		input.PlayedAt = now.UnixMilli()
	}
	if err := s.store.CreatePlayRecord(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *Service) GetPlayRecord(id string) (*model.PlayRecord, error) {
	return s.store.GetPlayRecord(id)
}

func (s *Service) ListPlayRecords(filter model.PlayRecordFilter, page, size int) ([]*model.PlayRecord, int, error) {
	all := s.store.ListPlayRecords()
	matched := make([]*model.PlayRecord, 0, len(all))
	for _, x := range all {
		if filter.Match(x) {
			matched = append(matched, x)
		}
	}
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].PlayedAt > matched[j].PlayedAt
	})
	total := len(matched)
	start := (page - 1) * size
	if start >= total {
		return []*model.PlayRecord{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return matched[start:end], total, nil
}

func (s *Service) DeletePlayRecord(id string) error {
	return s.store.DeletePlayRecord(id)
}
