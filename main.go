package main

import (
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup
var commandQueue sync.WaitGroup

var isShutdown bool

const (
	InfoColor    = "\033[1;34m%s\033[0m\n"
	NoticeColor  = "\033[1;36m%s\033[0m\n"
	WarningColor = "\033[1;33m%s\033[0m\n"
	ErrorColor   = "\033[1;31m%s\033[0m\n"
	DebugColor   = "\033[0;36m%s\033[0m\n"
)

func main() {
	fmt.Printf(DebugColor, "=================================================")
	fmt.Printf(DebugColor, "Loading application by aytronn")
	fmt.Printf(DebugColor, "Command \"help\" for see all command")
	fmt.Printf(DebugColor, "=================================================")

	onLoad()

	for {
		commandQueue.Add(1)
		go commandManager()
		commandQueue.Wait()
	}
}

func onLoad()  {
	wg.Add(1)
	go getConfig()
	wg.Wait()

	wg.Add(1)
	go getBetaUser()
	wg.Wait()


	getMongo("translations")


	registerCommand(Command{
		name: "search",
		Execute: searchCommand,
		usage: "search <pseudo> : allow to check if is in beta whitelist"})
	registerCommand(Command{
		name: "reload",
		Execute: reloadCommand,
		usage: "reload : Allow to reload"})
}

func shutdown()  {
	isShutdown = true

	os.Exit(0)
}



