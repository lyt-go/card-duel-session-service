// Card 领域模型（卡牌）。
package model

import (
	"strings"
	"time"
)

const (
	CardRarityCommon    = "common"
	CardRarityRare      = "rare"
	CardRarityEpic      = "epic"
	CardRarityLegendary = "legendary"

	CardTypeUnit      = "unit"
	CardTypeSpell     = "spell"
	CardTypeEquipment = "equipment"
)

type Card struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Cost        int       `json:"cost"`
	Attack      int       `json:"attack"`
	Health      int       `json:"health"`
	Rarity      string    `json:"rarity"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (x *Card) Validate() error {
	x.Name = strings.TrimSpace(x.Name)
	if x.Name == "" {
		return NewValidationError("name", "名称不能为空")
	}
	x.Description = strings.TrimSpace(x.Description)
	if x.Cost < 0 {
		return NewValidationError("cost", "费用不能为负数")
	}
	if x.Attack < 0 {
		return NewValidationError("attack", "攻击力不能为负数")
	}
	if x.Health < 0 {
		return NewValidationError("health", "生命值不能为负数")
	}
	if x.Rarity == "" {
		x.Rarity = CardRarityCommon
	}
	if !CardValidRarity(x.Rarity) {
		return NewValidationError("rarity", "稀有度不合法")
	}
	if x.Type == "" {
		x.Type = CardTypeUnit
	}
	if !CardValidType(x.Type) {
		return NewValidationError("type", "卡牌类型不合法")
	}
	return nil
}

func CardValidRarity(s string) bool {
	switch s {
	case CardRarityCommon, CardRarityRare, CardRarityEpic, CardRarityLegendary:
		return true
	default:
		return false
	}
}

func CardValidType(s string) bool {
	switch s {
	case CardTypeUnit, CardTypeSpell, CardTypeEquipment:
		return true
	default:
		return false
	}
}

type CardFilter struct {
	Rarity  string
	Type    string
	Keyword string
}

func (f CardFilter) Match(x *Card) bool {
	if f.Rarity != "" && x.Rarity != f.Rarity {
		return false
	}
	if f.Type != "" && x.Type != f.Type {
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
