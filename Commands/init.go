package commands

import (
	"fmt"

	config "github.com/Chris-Kellett/uma-discord-bot/Config"
	logger "github.com/Chris-Kellett/uma-discord-bot/Logger"
	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Command() *discordgo.ApplicationCommand
	Handler(*discordgo.InteractionCreate, string)
}

var (
	Commands = make(map[string]Command)
)

func Init() bool {
	Commands[UmaInfo{}.Command().Name] = UmaInfo{}
	err := resetAndRegisterCommands()
	return err == nil
}

func resetAndRegisterCommands() error {

	// Fetch existing global application commands
	appId := config.DISCORD_SESSION.State.User.ID
	existingCommands, err := config.DISCORD_SESSION.ApplicationCommands(appId, "")
	if err != nil {
		logger.ErrorText("COMMANDS", fmt.Sprintf("failed to get existing commands: %v", err))
		return err
	}

	// Delete all existing commands
	for _, cmd := range existingCommands {
		err := config.DISCORD_SESSION.ApplicationCommandDelete(appId, "", cmd.ID)
		if err != nil {
			logger.ErrorText("COMMANDS", fmt.Sprintf("failed to delete command %s: %v", cmd.Name, err))
		} else {
			logger.Info("COMMANDS", fmt.Sprintf("deleted command: %s", cmd.Name))
		}
	}

	// Register new commands from the Commands map
	for name, command := range Commands {
		_, err := config.DISCORD_SESSION.ApplicationCommandCreate(appId, "", command.Command())
		if err != nil {
			logger.ErrorText("COMMANDS", fmt.Sprintf("failed to register command %s: %v", name, err))
		} else {
			logger.Info("COMMANDS", fmt.Sprintf("registered command: %s", name))
		}
	}

	return nil
}
