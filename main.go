package main

import (
	"fmt"

	"github.com/cesardelgadom/logger/logger"
	pkgtest "github.com/cesardelgadom/logger/pkgTest"
)

func main() {
	logger.InitLooger("main", "", "NameProcess", true)
	logger.Info("Procesando...")
	logger.SetNamePkg("main")
	fmt.Println(logger.GetNamePkg())
	pkgtest.TestLogger()
	logger.Alert("Pilas mi perro")
	fmt.Println(logger.GetNamePkg())
	logger.CloseLog()
}
