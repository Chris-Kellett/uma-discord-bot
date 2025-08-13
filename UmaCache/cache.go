package umacache

import (
	"sync"

	datasets "github.com/Chris-Kellett/uma-discord-bot/Datasets"
)

// Card Data
var (
	cardData   map[int]datasets.CardData = make(map[int]datasets.CardData)
	cardDataMu sync.RWMutex
)

func GetCardData(cardId int) (datasets.CardData, bool) {
	cardDataMu.RLock()
	val, ok := cardData[cardId]
	cardDataMu.RUnlock()
	return val, ok
}

func SetCardData(data datasets.CardData) {
	cardDataMu.Lock()
	cardData[data.ID] = data
	cardDataMu.Unlock()
}
