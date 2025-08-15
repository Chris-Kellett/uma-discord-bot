package commands

import (
	"fmt"
	"strings"

	requests "github.com/Chris-Kellett/uma-discord-bot/Cache/Requests"
	umacache "github.com/Chris-Kellett/uma-discord-bot/Cache/Uma"
	config "github.com/Chris-Kellett/uma-discord-bot/Config"
	helpers "github.com/Chris-Kellett/uma-discord-bot/Helpers"
	logger "github.com/Chris-Kellett/uma-discord-bot/Logger"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

type UmaInfo struct{}

func (obj UmaInfo) Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "uma-info",
		Description: "See information about a specific Uma",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Search for your Uma, part-phrases are accepted",
				Required:    true,
			},
		},
	}
}

func (obj UmaInfo) Handler(i *discordgo.InteractionCreate, correlationId string) {
	logger.DebugWithCID(i.GuildID, correlationId, "UmaInfo Handler start")

	// Get the Request
	request, exists := requests.Get(correlationId)
	if !exists {
		helpers.SendError(i, "Unable to find Request data")
		requests.Complete(correlationId)
		return
	}

	// Validate the Input
	searchTerm, exists := request.Values.String["name"]
	if !exists {
		helpers.SendError(i, "Enter the name of an Uma to search for")
		requests.Complete(correlationId)
		return
	}
	searchTerm = strings.TrimSpace(searchTerm)

	if len(searchTerm) < 3 {
		helpers.SendError(i, "Please enter at least 3 search characters")
		requests.Complete(correlationId)
		return
	}

	// Search for the Uma
	char, found := umacache.SearchForCharacter(searchTerm)
	if !found {
		helpers.SendError(i, fmt.Sprintf("No Uma's found for the search term: '%s'", searchTerm))
		requests.Complete(correlationId)
		return
	}

	// Return the Uma
	e := embed.NewEmbed()
	e.SetTitle(char.Name)
	e.SetThumbnail(char.Info.ThumbImg)
	e.SetDescription(char.Info.TailFact)
	e.SetColor(config.EmbedColourGreen)
	helpers.SendEmbed(i, e)
}
