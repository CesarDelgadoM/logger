package message

import (
	"fmt"

	"github.com/cesardelgadom/logger/utils"
)

const (
	Info  = "[INFO] - "
	Alert = "[ALERT] - "
	Err   = "[ERROR] - "
	Fatal = "[FATAL] - "
)

func PrepareMsg(prefix string, msg string) string {
	//return fmt.Sprint(prefix, utils.GetTime(), " - ", namePkg, "] -> ", msg)
	return fmt.Sprint(prefix, utils.GetTime(), msg)
}
