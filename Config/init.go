package config

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var (
	// APP_x - relate to the current running instance of the bot
	APP_ROOTDIR      string
	APP_FOLDERNAME   string
	APP_HOSTNAME     string
	APP_SESSIONID    string
	APP_ISDEV        bool
	APP_LOGGINGLEVEL int

	// DISCORD_x - relate to interaction with discord
	DISCORD_BOTTOKEN string
	DISCORD_SESSION  *discordgo.Session

	// DB_x - relate to the database
	DB_IPADDRESS string
	DB_NAME      string
	DB_USER      string
	DB_PASSWORD  string
	DB_PORT      string
)

func Init() {
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		fmt.Printf("wdErr ERROR: %s\n", wdErr.Error())
		panic(wdErr)
	}

	trimIndex := strings.LastIndex(wd, `uma-discord-bot\`)
	if trimIndex != -1 {
		APP_ROOTDIR = wd[:trimIndex+len(`uma-discord-bot\`)]
	} else {
		APP_ROOTDIR = wd
	}

	loadErr := godotenv.Load(path.Join(APP_ROOTDIR, ".env"))
	if loadErr != nil {
		fmt.Printf("loadErr ERROR: %s\n", loadErr.Error())
		panic(loadErr)
	}

	envErr := parseEnvVariables()
	if envErr != nil {
		fmt.Printf("parseEnvVariables ERROR: %s\n", envErr.Error())
		panic(envErr)
	}

	currentHostName, err := os.Hostname()
	if err != nil {
		APP_HOSTNAME = "Unknown"
	} else {
		APP_HOSTNAME = currentHostName
	}

	APP_SESSIONID = uuid.New().String()
}

func parseEnvVariables() error {
	var (
		err               error
		value             string
		notFoundText      string = "could not find .env value: "
		skippingText      string = "optional .env value not provided: "
		optionalErrorText string = "error processing optional .env value: "
	)

	value = os.Getenv("APP_ISDEV")
	if value == "" {
		return fmt.Errorf("%s%s", notFoundText, "APP_ISDEV")
	} else {
		APP_ISDEV, err = strconv.ParseBool(value)
		if err != nil {
			return err
		}
	}

	APP_FOLDERNAME = os.Getenv("APP_FOLDERNAME")
	if APP_FOLDERNAME == "" {
		return fmt.Errorf("%s%s", notFoundText, "APP_FOLDERNAME")
	}

	DISCORD_BOTTOKEN = os.Getenv("DISCORD_BOTTOKEN")
	if DISCORD_BOTTOKEN == "" {
		return fmt.Errorf("%s%s", notFoundText, "DISCORD_BOTTOKEN")
	}

	DB_IPADDRESS = os.Getenv("DB_IPADDRESS")
	if DB_IPADDRESS == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_IPADDRESS")
	}

	DB_NAME = os.Getenv("DB_NAME")
	if DB_NAME == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_NAME")
	}

	DB_USER = os.Getenv("DB_USER")
	if DB_USER == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_USER")
	}

	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_PASSWORD")
	}

	DB_PORT = os.Getenv("DB_PORT")
	if DB_PORT == "" {
		return fmt.Errorf("%s%s", notFoundText, "DB_PORT")
	}

	value = os.Getenv("APP_LOGGINGLEVEL")
	if value == "" {
		fmt.Printf("%s%s\n", skippingText, "APP_LOGGINGLEVEL")
	} else {
		APP_LOGGINGLEVEL, err = strconv.Atoi(value)
		if err != nil {
			fmt.Printf("%s%s :: %s\n", optionalErrorText, "APP_LOGGINGLEVEL", err)
		}
	}

	return nil
}
