package umacache

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	datasets "github.com/Chris-Kellett/uma-discord-bot/Datasets"
	helpers "github.com/Chris-Kellett/uma-discord-bot/Helpers"
	logger "github.com/Chris-Kellett/uma-discord-bot/Logger"
)

func getCharacters() {
	funcName := "Characters"
	delete(activeErrors, funcName)

	body, err := helpers.GetBytesFromURL("UMACACHE", tracCharacters)
	if err != nil {
		activeErrors[funcName] = fmt.Sprintf("Unable to obtain %s data", funcName)
		return
	}
	var list []datasets.Character
	if err := json.Unmarshal(body, &list); err != nil {
		logger.Error("UMACACHE", err)
		activeErrors[funcName] = fmt.Sprintf("Unable to process %s data", funcName)
		return
	}

	logger.Debug("UMACACHE", "%s: Adding to cache, %d entries", funcName, len(list))
	for _, data := range list {
		charactersMu.Lock()
		characters[data.ID] = data
		charactersMu.Unlock()
	}
	logger.Event("UMACACHE", "%s: Cache populated, %d entries", funcName, len(list))
}

func getCharacterData() {
	funcName := "Character Data"
	delete(activeErrors, funcName)

	body, err := helpers.GetBytesFromURL("UMACACHE", tracCharacterData)
	if err != nil {
		activeErrors[funcName] = fmt.Sprintf("Unable to obtain %s data", funcName)
		return
	}
	var list []datasets.CharacterStats
	if err := json.Unmarshal(body, &list); err != nil {
		logger.Error("UMACACHE", err)
		activeErrors[funcName] = fmt.Sprintf("Unable to process %s data", funcName)
		return
	}

	logger.Debug("UMACACHE", "%s: Adding to cache, %d entries", funcName, len(list))

	// Need to implement this ticker as Umap API has a max of 10 requests a second
	ticker := time.NewTicker(110 * time.Millisecond)
	defer ticker.Stop()

	success := 0
	for _, data := range list {
		<-ticker.C

		infoUrl := umapCharacterInfo + strconv.Itoa(data.ID)
		infoBody, err := helpers.GetBytesFromURL("UMACACHE", infoUrl)
		if err != nil {
			logger.ErrorText("UMACACHE", "Character ID %d: %s", data.ID, err)
			return
		}
		var info datasets.CharacterInfo
		if err := json.Unmarshal(infoBody, &info); err != nil {
			logger.ErrorText("UMACACHE", "Character ID %d: %s", data.ID, err)
			continue
		}
		charactersMu.Lock()
		char, ok := characters[data.ID]
		if !ok {
			charactersMu.Unlock()
			continue
		}
		char.Stats = data
		char.Info = info
		char.Name = info.NameEn
		characters[data.ID] = char
		charactersMu.Unlock()
		success++
	}
	logger.Event("UMACACHE", "%s: Cache populated, %d entries", funcName, success)
}
