package umacache

import (
	"strings"
	"sync"

	datasets "github.com/Chris-Kellett/uma-discord-bot/Datasets"
)

// Characters
var (
	characters   map[int]datasets.Character = make(map[int]datasets.Character)
	charactersMu sync.RWMutex
)

func SearchForCharacter(searchTerm string) (datasets.Character, bool) {
	searchTerm = strings.ToLower(searchTerm)
	var character datasets.Character
	found := false
	charactersMu.RLock()
	for _, char := range characters {
		if strings.Contains(strings.ToLower(char.Name), searchTerm) {
			character = char
			found = true
			break
		}
	}
	charactersMu.RUnlock()
	return character, found
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
