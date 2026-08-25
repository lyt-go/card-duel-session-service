package service

// Overview 返回全局统计总览：每个实体的总数与按状态分布。
func (s *Service) Overview() map[string]interface{} {
	return map[string]interface{}{
		"cards":        s.countCard(),
		"decks":        s.countDeck(),
		"games":        s.countGame(),
		"play-records": s.countPlayRecord(),
		"players":      s.countPlayer(),
		"collections":  s.countCollection(),
	}
}

func (s *Service) countCard() map[string]interface{} {
	items := s.store.ListCards()
	m := map[string]int{}
	for _, it := range items {
		m[it.Rarity]++
	}
	return map[string]interface{}{"total": len(items), "by_rarity": m}
}

func (s *Service) countDeck() map[string]interface{} {
	items := s.store.ListDecks()
	m := map[string]int{}
	for _, it := range items {
		m[it.Status]++
	}
	return map[string]interface{}{"total": len(items), "by_status": m}
}

func (s *Service) countGame() map[string]interface{} {
	items := s.store.ListGames()
	m := map[string]int{}
	for _, it := range items {
		m[it.Status]++
	}
	return map[string]interface{}{"total": len(items), "by_status": m}
}

func (s *Service) countPlayRecord() map[string]interface{} {
	items := s.store.ListPlayRecords()
	return map[string]interface{}{"total": len(items)}
}

func (s *Service) countPlayer() map[string]interface{} {
	items := s.store.ListPlayers()
	m := map[string]int{}
	for _, it := range items {
		m[it.Status]++
	}
	return map[string]interface{}{"total": len(items), "by_status": m}
}

func (s *Service) countCollection() map[string]interface{} {
	items := s.store.ListCollections()
	return map[string]interface{}{"total": len(items)}
}
