// Deck 领域模型（牌组）。
package model

import (
	"strings"
	"time"
)

const (
	DeckStatusDraft    = "draft"
	DeckStatusActive   = "active"
	DeckStatusArchived = "archived"
)

type Deck struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CardCount   int       `json:"card_count"`
	PlayerID    string    `json:"player_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (x *Deck) Validate() error {
	x.Name = strings.TrimSpace(x.Name)
	if x.Name == "" {
		return NewValidationError("name", "名称不能为空")
	}
	x.Description = strings.TrimSpace(x.Description)
	if x.CardCount < 0 {
		return NewValidationError("card_count", "卡牌数量不能为负数")
	}
	x.PlayerID = strings.TrimSpace(x.PlayerID)
	if x.Status == "" {
		x.Status = DeckStatusDraft
	}
	if !DeckValidStatus(x.Status) {
		return NewValidationError("status", "状态不合法")
	}
	return nil
}

func DeckValidStatus(s string) bool {
	switch s {
	case DeckStatusDraft, DeckStatusActive, DeckStatusArchived:
		return true
	default:
		return false
	}
}

var deckTransitions = map[string]map[string]bool{
	DeckStatusDraft:    {DeckStatusActive: true, DeckStatusArchived: true},
	DeckStatusActive:   {DeckStatusArchived: true},
	DeckStatusArchived: {DeckStatusActive: true},
}

func DeckCanTransition(from, to string) bool {
	if m, ok := deckTransitions[from]; ok {
		return m[to]
	}
	return false
}

type DeckFilter struct {
	Status   string
	PlayerID string
	Keyword  string
}

func (f DeckFilter) Match(x *Deck) bool {
	if f.Status != "" && x.Status != f.Status {
		return false
	}
	if f.PlayerID != "" && x.PlayerID != f.PlayerID {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(x.Name), k) &&
			!strings.Contains(strings.ToLower(x.Description), k) {
			return false
		}
	}
	return true
}
