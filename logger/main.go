package main

import (
	"logger/lib/logger"
)


func main() {
	logger.Info("Logger, Info!")
	logger.Error("Logger, Error!", nil)
	logger.Debug("Logger, Debug!")
	logger.Warn("Logger, Warn!")
	logger.Fatal("Logger, Fatal!")
	logger.Panic("Logger, Logger!")
}
