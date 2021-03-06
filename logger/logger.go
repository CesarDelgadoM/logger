package logger

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cesardelgadom/logger/message"
)

type logger struct {
	PathFile   string
	NameFile   string
	PrintTrace bool
}

var log *logger
var writer *bufio.Writer

func InitLooger(pathFile string, nameFile string, printTrace bool) {
	CreateFile(pathFile + nameFile)
	writer = GetWriter()
	log = &logger{
		PathFile:   pathFile,
		NameFile:   nameFile,
		PrintTrace: printTrace,
	}
}

func (log *logger) logger(msg string) {
	if log != nil {
		if log.PrintTrace {
			fmt.Println(msg)
		}
		_, err := writer.WriteString(msg + "\n")
		if err != nil {
			Error(err.Error())
		}
		writer.Flush()
	} else {
		fmt.Println("Logger not inicializate")
		os.Exit(1)
	}
}

func Info(msg string) {
	msgInfo := message.PrepareMsg(message.Info, msg)
	log.logger(msgInfo)
}

func Alert(msg string) {
	msgAlert := message.PrepareMsg(message.Alert, msg)
	log.logger(msgAlert)
}

func Error(msg string) {
	msgError := message.PrepareMsg(message.Err, msg)
	log.logger(msgError)
	os.Exit(1)
}

func Fatal(msg string) {
	msgFatal := message.PrepareMsg(message.Fatal, msg)
	log.logger(msgFatal)
	os.Exit(1)
}
