package logger

import (
	"bytes"
	glog "github.com/kataras/golog"
	"log"
	"net/http"
	"os"
)

var Logger = glog.New()

type remote struct{}

//Write implementation to send to remote server
func (r remote) Write(data []byte) (n int, err error) {
	go func() {
		req, err := http.NewRequest("POST",
			"http://localhost:8010/log",
			bytes.NewBuffer(data),
		)

		if err == nil {
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, _ := client.Do(req)
			defer resp.Body.Close()
		}
	}()
	return len(data), nil
}

//SetLoggingOutput
func SetLoggingOutput(localStdout bool) {
	if localStdout {
		configureLocal()
		return
	}
	configureRemote()
}

//configureLocal for local implementation
func configureLocal() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Logger.SetOutput(os.Stdout)
	Logger.SetLevel("debug")
	Logger.SetLevelOutput("info", file)
}

//configureRemote for remote logger configuration
func configureRemote() {
	r := remote{}
	Logger.SetLevelFormat("info", "json")
	Logger.SetLevelOutput("info", r)
}
