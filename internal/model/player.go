// Player 领域模型（玩家）。
package model

import (
	"strings"
	"time"
)

const (
	PlayerStatusOffline = "offline"
	PlayerStatusOnline  = "online"
)

type Player struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Level     int       `json:"level"`
	Score     int       `json:"score"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (x *Player) Validate() error {
	x.Nickname = strings.TrimSpace(x.Nickname)
	if x.Nickname == "" {
		return NewValidationError("nickname", "昵称不能为空")
	}
	if x.Level <= 0 {
		return NewValidationError("level", "等级必须为正数")
	}
	if x.Score < 0 {
		return NewValidationError("score", "积分不能为负数")
	}
	if x.Status == "" {
		x.Status = PlayerStatusOffline
	}
	if !PlayerValidStatus(x.Status) {
		return NewValidationError("status", "状态不合法")
	}
	return nil
}

func PlayerValidStatus(s string) bool {
	switch s {
	case PlayerStatusOffline, PlayerStatusOnline:
		return true
	default:
		return false
	}
}

var playerTransitions = map[string]map[string]bool{
	PlayerStatusOffline: {PlayerStatusOnline: true},
	PlayerStatusOnline:  {PlayerStatusOffline: true},
}

func PlayerCanTransition(from, to string) bool {
	if m, ok := playerTransitions[from]; ok {
		return m[to]
	}
	return false
}

type PlayerFilter struct {
	Status  string
	Keyword string
}

func (f PlayerFilter) Match(x *Player) bool {
	if f.Status != "" && x.Status != f.Status {
		return false
	}
	if f.Keyword != "" {
		k := strings.ToLower(strings.TrimSpace(f.Keyword))
		if k != "" && !strings.Contains(strings.ToLower(x.Nickname), k) {
			return false
		}
	}
	return true
}
