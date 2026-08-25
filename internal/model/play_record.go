// PlayRecord 领域模型（出牌记录）。
package model

import (
	"strings"
	"time"
)

type PlayRecord struct {
	ID        string    `json:"id"`
	GameID    string    `json:"game_id"`
	DeckID    string    `json:"deck_id"`
	CardID    string    `json:"card_id"`
	Turn      int       `json:"turn"`
	PlayedAt  int64     `json:"played_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (x *PlayRecord) Validate() error {
	x.GameID = strings.TrimSpace(x.GameID)
	if x.GameID == "" {
		return NewValidationError("game_id", "对局 ID 不能为空")
	}
	x.DeckID = strings.TrimSpace(x.DeckID)
	if x.DeckID == "" {
		return NewValidationError("deck_id", "牌组 ID 不能为空")
	}
	x.CardID = strings.TrimSpace(x.CardID)
	if x.CardID == "" {
		return NewValidationError("card_id", "卡牌 ID 不能为空")
	}
	if x.Turn < 0 {
		return NewValidationError("turn", "回合数不能为负数")
	}
	return nil
}

type PlayRecordFilter struct {
	GameID string
	DeckID string
}

func (f PlayRecordFilter) Match(x *PlayRecord) bool {
	if f.GameID != "" && x.GameID != f.GameID {
		return false
	}
	if f.DeckID != "" && x.DeckID != f.DeckID {
		return false
	}
	return true
}
