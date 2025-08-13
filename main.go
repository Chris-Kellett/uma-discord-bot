package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	config "github.com/Chris-Kellett/uma-discord-bot/Config"
	logger "github.com/Chris-Kellett/uma-discord-bot/Logger"
)

func main() {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	config.Init()
	logger.Init()
	logger.Event("CONFIG", "Config initialisation successful, Session ID: %s", config.APP_SESSIONID)
	<-osSignal
	stop()
}
func stop() {
	logger.Info("MAIN", "OS Quit signal received, bot stopping...")

	fmt.Println("Bot gracefully stopped...")
	os.Exit(1)
}
