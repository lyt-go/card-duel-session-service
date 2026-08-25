// Game 领域模型（对局）。
package model

import (
	"strings"
	"time"
)

const (
	GameStatusWaiting  = "waiting"
	GameStatusPlaying  = "playing"
	GameStatusFinished = "finished"
)

type Game struct {
	ID        string    `json:"id"`
	DeckAID   string    `json:"deck_a_id"`
	DeckBID   string    `json:"deck_b_id"`
	Turn      int       `json:"turn"`
	WinnerID  string    `json:"winner_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (x *Game) Validate() error {
	x.DeckAID = strings.TrimSpace(x.DeckAID)
	if x.DeckAID == "" {
		return NewValidationError("deck_a_id", "甲方牌组 ID 不能为空")
	}
	x.DeckBID = strings.TrimSpace(x.DeckBID)
	if x.DeckBID == "" {
		return NewValidationError("deck_b_id", "乙方牌组 ID 不能为空")
	}
	if x.DeckAID == x.DeckBID {
		return NewValidationError("deck_b_id", "双方牌组不能相同")
	}
	if x.Turn < 0 {
		return NewValidationError("turn", "回合数不能为负数")
	}
	x.WinnerID = strings.TrimSpace(x.WinnerID)
	if x.Status == "" {
		x.Status = GameStatusWaiting
	}
	if !GameValidStatus(x.Status) {
		return NewValidationError("status", "状态不合法")
	}
	return nil
}

func GameValidStatus(s string) bool {
	switch s {
	case GameStatusWaiting, GameStatusPlaying, GameStatusFinished:
		return true
	default:
		return false
	}
}

var gameTransitions = map[string]map[string]bool{
	GameStatusWaiting: {GameStatusPlaying: true},
	GameStatusPlaying: {GameStatusFinished: true},
}

func GameCanTransition(from, to string) bool {
	if m, ok := gameTransitions[from]; ok {
		return m[to]
	}
	return false
}

type GameFilter struct {
	Status  string
	DeckAID string
}

func (f GameFilter) Match(x *Game) bool {
	if f.Status != "" && x.Status != f.Status {
		return false
	}
	if f.DeckAID != "" && x.DeckAID != f.DeckAID {
		return false
	}
	return true
}
