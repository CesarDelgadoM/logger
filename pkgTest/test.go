package pkgtest

import (
	"github.com/cesardelgadom/logger/logger"
)

func TestLogger() {
	logger.SetNamePkg("TestLogger")
	test2Logger()
}

func test2Logger() {
	logger.Info("Estoy escribiendo desde otro paquete")
}
