package service

import (
	"errors"
	"fmt"

	"cardgame/internal/blueprint"
)

type BlueprintBuilder struct{ cache *blueprint.Cache }

func NewBlueprintBuilder(cache *blueprint.Cache) *BlueprintBuilder {
	return &BlueprintBuilder{cache: cache}
}

func (b *BlueprintBuilder) Build(id string, cards []string) (result *blueprint.Blueprint, err error) {
	result = &blueprint.Blueprint{ID: id}
	b.cache.Put(id, result)
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("蓝图构建失败: %v", recovered)
		}
	}()
	for _, card := range cards {
		result.Cards = append(result.Cards, card)
		if card == "panic-card" {
			panic("卡牌规则缺失")
		}
	}
	result.Rules = &blueprint.Rules{Allowed: map[string]bool{"ranked": true}}
	return result, nil
}

func (b *BlueprintBuilder) Serve(id string) (int, error) {
	item, ok := b.cache.Get(id)
	if !ok {
		return 0, errors.New("蓝图不存在")
	}
	return len(item.Rules.Allowed), nil
}
