package umacache

import (
	"encoding/json"

	datasets "github.com/Chris-Kellett/uma-discord-bot/Datasets"
	helpers "github.com/Chris-Kellett/uma-discord-bot/Helpers"
	logger "github.com/Chris-Kellett/uma-discord-bot/Logger"
)

var (
	// Errors
	activeErrors map[string]string = make(map[string]string)

	// API URLs
	urlCardData = "https://www.tracenacademy.com/api/CardData"
)

func Init() {
	getCardData()
}

func getCardData() {
	funcName := "Card Data"
	delete(activeErrors, funcName)

	body, err := helpers.GetBytesFromURL("UMACACHE", urlCardData)
	if err != nil {
		activeErrors[funcName] = "Unable to obtain global card data"
		return
	}
	var list []datasets.CardData
	if err := json.Unmarshal(body, &list); err != nil {
		logger.Error("UMACACHE", err)
		activeErrors[funcName] = "Unable to process global card data"
		return
	}

	logger.Debug("UMACACHE", "CardData: Adding to cache, %d entries", len(list))
	for _, data := range list {
		SetCardData(data)
	}
	logger.Debug("UMACACHE", "CardData: Cache populated")
}
