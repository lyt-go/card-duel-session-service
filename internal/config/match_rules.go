package config

import "errors"

type DeckRule interface {
	Validate(cardCount int) error
}

type StandardDeckRule struct {
	Minimum int
}

func (r *StandardDeckRule) Validate(cardCount int) error {
	if r == nil {
		return nil
	}
	if cardCount < r.Minimum {
		return errors.New("牌组卡牌数量不足")
	}
	return nil
}

type MatchUsage struct {
	Starts map[string]int
}

func LoadDeckRule(mode string) DeckRule {
	if mode == "standard" {
		return &StandardDeckRule{Minimum: 30}
	}
	var missing *StandardDeckRule
	return missing
}

func LoadMatchUsage(mode string) *MatchUsage {
	if mode == "standard" {
		return &MatchUsage{Starts: make(map[string]int)}
	}
	return &MatchUsage{}
}
