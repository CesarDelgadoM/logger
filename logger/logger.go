package logger

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cesardelgadom/logger/message"
)

type LoggerMethod interface {
	InitLooger()
	log(msg string)
}

type Logger struct {
	PathFile   string
	NameFile   string
	PrintTrace bool
}

var logger *Logger
var bufWrite *bufio.Writer
var file *os.File

func (log *Logger) InitLooger() {
	logger = log
	bufWrite = createFile(log.NameFile)
}

func (log *Logger) log(msg string) {
	if log != nil {
		if log.PrintTrace {
			fmt.Println(msg)
		}
		bufWrite.WriteString(msg)
	} else {
		fmt.Println("Logger not inicializate")
		os.Exit(1)
	}
}

func Info(msg string) {
	msgInfo := message.PrepareMsg(message.Info, msg)
	logger.log(msgInfo)
}

func Alert(msg string) {
	msgAlert := message.PrepareMsg(message.Alert, msg)
	logger.log(msgAlert)
}

func Error(msg string) {
	msgError := message.PrepareMsg(message.Err, msg)
	logger.log(msgError)
	os.Exit(1)
}

func Fatal(msg string) {
	msgFatal := message.PrepareMsg(message.Fatal, msg)
	logger.log(msgFatal)
	os.Exit(1)
}

func createFile(path string) *bufio.Writer {
	f, err := os.Create(path + ".log")
	isError(err)
	if !existFile(path) {
		//createFile(path)
	}
	bw := bufio.NewWriter(f)
	file = f
	return bw
}

func existFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	Info(fileInfo.Name())
	return true
}

func CloseLog() {
	file.Close()
}

func isError(err error) {
	if err != nil {
		Fatal(err.Error())
	}
}
