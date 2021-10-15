package main

import (
	"github.com/cesardelgadom/logger/logger"
)

func main() {
	log := logger.Logger{
		PathFile:   "",
		NameFile:   "NameProcess",
		PrintTrace: true,
	}
	log.InitLooger()
	logger.Info("Message of information...")
	logger.CloseLog()
}
