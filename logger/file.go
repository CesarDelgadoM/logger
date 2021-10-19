package logger

import (
	"bufio"
	"os"
)

var file *os.File

func CreateFile(path string) {
	f, err := os.Create(path + ".log")
	if err != nil {
		Fatal(err.Error())
	}
	file = f
}

func GetWriter() *bufio.Writer {
	return bufio.NewWriter(file)
}

func CloseLog() {
	file.Close()
}
