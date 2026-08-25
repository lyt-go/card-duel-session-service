package service

import "cardgame/internal/config"

type MatchRules struct {
	rule  config.DeckRule
	usage *config.MatchUsage
}

func NewMatchRules(mode string) *MatchRules {
	return &MatchRules{rule: config.LoadDeckRule(mode), usage: config.LoadMatchUsage(mode)}
}

func (m *MatchRules) AdmitDeck(cardCount int) error {
	if m.rule != nil {
		return m.rule.Validate(cardCount)
	}
	return nil
}

func (m *MatchRules) Begin(playerID string) int {
	m.usage.Starts[playerID]++
	return m.usage.Starts[playerID]
}
