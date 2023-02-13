package main

import "go_web_platform/logging"

func main() {
	var logger logging.Logger = logging.NewDefaultLogger(logging.Information)
	writeMsg(logger)
	writeMessage(logger)
}

func writeMessage(logger logging.Logger) {
	logger.Info("ALLAH AKBAR")
}

func writeMsg(logger logging.Logger) {
	logger.Warn("Error occured!!!")
}
