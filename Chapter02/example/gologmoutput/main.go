package main

import (
	"github.com/kataras/golog"
	"log"
	"os"
)

const logFile = "infolog.txt"

func init() {
	golog.SetLevel("debug")
	configureLogger()
}

func main() {
	golog.Println("This is a raw message, no levels, no colors.")
	golog.Info("This is an info message, with colors (if the output is terminal)")
	golog.Warn("This is a warning message")
	golog.Error("This is an error message")
	golog.Debug("This is a debug message")
	golog.Fatal(`Fatal will exit no matter what`)
}

func configureLogger() {
	// open infolog.txt  append if exist (os.O_APPEND) or create if not (os.O_CREATE) and read write (os.O_WRONLY)
	infof, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	golog.SetLevelOutput("info", infof)

	// open infoerr.txt  append if exist (os.O_APPEND) or create if not (os.O_CREATE) and read write (os.O_WRONLY)
	errf, err := os.OpenFile("infoerr.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	golog.SetLevelOutput("error", errf)
}
