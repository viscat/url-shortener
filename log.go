package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"os"
	"path/filepath"
)

// Initializes log
func initLogger(verbosity bool) {
	if verbosity {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	log.SetOutput(os.Stdout)
	formatter := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true}
	log.SetFormatter(formatter)

	hook := lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  createLogFile("info.log"),
		log.ErrorLevel: createLogFile("error.log"),
	})
	hook.SetFormatter(formatter)
	log.AddHook(hook)
}
func createLogFile(fileName string) string {
	infoLogFile := filepath.Join(os.Getenv("URLSHORTENER_LOG_DIR"), fileName)
	_, err := os.OpenFile(infoLogFile, os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Could not create log file %v", err)
	}
	return infoLogFile
}
