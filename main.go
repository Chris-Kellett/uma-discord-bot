package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	bot "github.com/Chris-Kellett/uma-discord-bot/Bot"
	config "github.com/Chris-Kellett/uma-discord-bot/Config"
	logger "github.com/Chris-Kellett/uma-discord-bot/Logger"
	ping "github.com/Chris-Kellett/uma-discord-bot/Ping"
	umacache "github.com/Chris-Kellett/uma-discord-bot/UmaCache"
)

func main() {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	config.Init()
	logger.Init()
	logger.Event("CONFIG", "Config initialisation successful, Session ID: %s", config.APP_SESSIONID)
	umacache.Init()
	go ping.Init()
	go bot.Init()
	<-osSignal
	stop()
}
func stop() {
	logger.Info("MAIN", "OS Quit signal received, bot stopping...")
	bot.Stop <- true
	logger.Stop <- true
	fmt.Println("Bot gracefully stopped...")
	os.Exit(0)
}
