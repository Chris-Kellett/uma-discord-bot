package umacache

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	datasets "github.com/Chris-Kellett/uma-discord-bot/Datasets"
	helpers "github.com/Chris-Kellett/uma-discord-bot/Helpers"
	logger "github.com/Chris-Kellett/uma-discord-bot/Logger"
)

func getSupportCards() {
	funcName := "Support Cards"
	delete(activeErrors, funcName)

	body, err := helpers.GetBytesFromURL("UMACACHE", tracSupportCards)
	if err != nil {
		activeErrors[funcName] = fmt.Sprintf("Unable to obtain %s data", funcName)
		return
	}
	var list []datasets.SupportCard
	if err := json.Unmarshal(body, &list); err != nil {
		logger.Error("UMACACHE", err)
		activeErrors[funcName] = fmt.Sprintf("Unable to process %s data", funcName)
		return
	}

	logger.Debug("UMACACHE", "%s: Adding to cache, %d entries", funcName, len(list))
	for _, data := range list {
		supportCardsMu.Lock()
		supportCards[data.ID] = data
		supportCardsMu.Unlock()
	}
	logger.Event("UMACACHE", "%s: Cache populated, %d entries", funcName, len(list))
}

func getSupportCardData() {
	funcName := "Support Card Data"
	delete(activeErrors, funcName)

	body, err := helpers.GetBytesFromURL("UMACACHE", tracSupportCardsData)
	if err != nil {
		activeErrors[funcName] = fmt.Sprintf("Unable to obtain %s data", funcName)
		return
	}
	var list []datasets.SupportCardData
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

		infoUrl := umapSupportCardInfo + strconv.Itoa(data.ID)
		infoBody, err := helpers.GetBytesFromURL("UMACACHE", infoUrl)
		if err != nil {
			logger.ErrorText("UMACACHE", "Support Card ID %d: %s", data.ID, err)
			return
		}
		var info datasets.SupportCardInfo
		if err := json.Unmarshal(infoBody, &info); err != nil {
			logger.ErrorText("UMACACHE", "Support Card ID %d: %s", data.ID, err)
			continue
		}
		supportCardsMu.Lock()
		card, ok := supportCards[data.ID]
		if !ok {
			supportCardsMu.Unlock()
			continue
		}
		card.Data = data
		card.Info = info
		card.Name = strings.Trim(info.TitleEN, "[]")
		supportCards[data.ID] = card
		supportCardsMu.Unlock()
		success++
	}
	logger.Event("UMACACHE", "%s: Cache populated, %d entries", funcName, success)
}
