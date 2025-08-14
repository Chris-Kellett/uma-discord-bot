package umacache

import (
	"encoding/csv"
	"fmt"
	"net/http"

	logger "github.com/Chris-Kellett/uma-discord-bot/Logger"
)

func getTranslationMap(url string) (map[string]string, error) {
	logger.Debug("UMACACHE", "Getting Translations from URL: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("UMACACHE", err)
		return nil, err
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	translations := make(map[string]string)
	lineNum := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Warning: skipping line %d: %v\n", lineNum+1, err)
			lineNum++
			continue
		}

		if lineNum == 0 {
			lineNum++
			continue // skip header
		}

		var key, value string

		switch len(record) {
		case 0:
			// skip empty line
			lineNum++
			continue
		case 1:
			key = record[0]
			value = key
		default:
			key = record[0]
			value = record[1]
			if value == "" {
				value = key
			}
		}

		translations[key] = value
		lineNum++
	}

	return translations, nil
}

func OutputCache() {
	for _, char := range characters {
		logger.Event("CACHE", "%v", char)
		logger.Debug("CACHE", "--------------------------------------------")
	}
}
