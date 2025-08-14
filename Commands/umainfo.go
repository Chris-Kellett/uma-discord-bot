package commands

import (
	logger "github.com/Chris-Kellett/uma-discord-bot/Logger"
	"github.com/bwmarrin/discordgo"
)

type UmaInfo struct{}

func (obj UmaInfo) Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "uma-info",
		Description: "See information about a specific Uma",
	}
}

func (obj UmaInfo) Handler(i *discordgo.InteractionCreate, correlationId string) {
	logger.DebugWithCID(i.GuildID, correlationId, "UmaInfo Handler start")
}
