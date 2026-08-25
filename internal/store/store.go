// Package store 定义数据访问接口与内存实现。
package store

import (
	"errors"

	"cardgame/internal/model"
)

var (
	ErrNotFound = errors.New("记录不存在")
	ErrConflict = errors.New("记录已存在或状态冲突")
)

// Store 聚合全部实体的数据访问方法，便于测试时替换实现。
type Store interface {
	CreateCard(x *model.Card) error
	GetCard(id string) (*model.Card, error)
	GetCardByName(v string) (*model.Card, error)
	ListCards() []*model.Card
	UpdateCard(x *model.Card) error
	DeleteCard(id string) error
	CreateDeck(x *model.Deck) error
	GetDeck(id string) (*model.Deck, error)
	GetDeckByName(v string) (*model.Deck, error)
	ListDecks() []*model.Deck
	UpdateDeck(x *model.Deck) error
	DeleteDeck(id string) error
	CreateGame(x *model.Game) error
	GetGame(id string) (*model.Game, error)
	ListGames() []*model.Game
	UpdateGame(x *model.Game) error
	DeleteGame(id string) error
	CreatePlayRecord(x *model.PlayRecord) error
	GetPlayRecord(id string) (*model.PlayRecord, error)
	ListPlayRecords() []*model.PlayRecord
	DeletePlayRecord(id string) error
	CreatePlayer(x *model.Player) error
	GetPlayer(id string) (*model.Player, error)
	GetPlayerByNickname(v string) (*model.Player, error)
	ListPlayers() []*model.Player
	UpdatePlayer(x *model.Player) error
	DeletePlayer(id string) error
	CreateCollection(x *model.Collection) error
	GetCollection(id string) (*model.Collection, error)
	ListCollections() []*model.Collection
	DeleteCollection(id string) error
}
