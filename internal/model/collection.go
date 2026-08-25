// Collection 领域模型（卡册/收藏）。
package model

import (
	"strings"
	"time"
)

type Collection struct {
	ID        string    `json:"id"`
	PlayerID  string    `json:"player_id"`
	CardID    string    `json:"card_id"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

func (x *Collection) Validate() error {
	x.PlayerID = strings.TrimSpace(x.PlayerID)
	if x.PlayerID == "" {
		return NewValidationError("player_id", "玩家 ID 不能为空")
	}
	x.CardID = strings.TrimSpace(x.CardID)
	if x.CardID == "" {
		return NewValidationError("card_id", "卡牌 ID 不能为空")
	}
	if x.Count <= 0 {
		return NewValidationError("count", "数量必须为正数")
	}
	return nil
}

type CollectionFilter struct {
	PlayerID string
	CardID   string
}

func (f CollectionFilter) Match(x *Collection) bool {
	if f.PlayerID != "" && x.PlayerID != f.PlayerID {
		return false
	}
	if f.CardID != "" && x.CardID != f.CardID {
		return false
	}
	return true
}
