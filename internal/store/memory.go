package store

import (
	"sync"

	"cardgame/internal/model"
)

type MemoryStore struct {
	mu          sync.RWMutex
	cards       map[string]*model.Card
	decks       map[string]*model.Deck
	games       map[string]*model.Game
	playRecords map[string]*model.PlayRecord
	players     map[string]*model.Player
	collections map[string]*model.Collection
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		cards:       make(map[string]*model.Card),
		decks:       make(map[string]*model.Deck),
		games:       make(map[string]*model.Game),
		playRecords: make(map[string]*model.PlayRecord),
		players:     make(map[string]*model.Player),
		collections: make(map[string]*model.Collection),
	}
}

var _ Store = (*MemoryStore)(nil)
