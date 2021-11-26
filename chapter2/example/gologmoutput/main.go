package main

import (
	"github.com/kataras/golog"
	"log"
	"os"
)

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
	infof, err := os.OpenFile("infolog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	golog.SetLevelOutput("info", infof)

	errf, err := os.OpenFile("infoerr.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	golog.SetLevelOutput("error", errf)
}
