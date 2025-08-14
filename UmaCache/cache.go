package umacache

import (
	"sync"

	datasets "github.com/Chris-Kellett/uma-discord-bot/Datasets"
)

// Characters
var (
	characters   map[int]datasets.Character = make(map[int]datasets.Character)
	charactersMu sync.RWMutex
)

func Character(characterId int) (datasets.Character, bool) {
	charactersMu.RLock()
	val, ok := characters[characterId]
	charactersMu.RUnlock()
	return val, ok
}

// Support Cards
var (
	supportCards   map[int]datasets.SupportCard = make(map[int]datasets.SupportCard)
	supportCardsMu sync.RWMutex
)

func Card(cardId int) (datasets.SupportCard, bool) {
	supportCardsMu.RLock()
	val, ok := supportCards[cardId]
	supportCardsMu.RUnlock()
	return val, ok
}
