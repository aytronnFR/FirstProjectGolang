package main

import "log"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(ErrorColor + ": %s", msg, err)
	}
}

func info(msg string) {
	log.Printf(WarningColor, msg)
}
