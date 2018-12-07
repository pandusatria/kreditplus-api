package commons

import (
	"fmt"
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
)

// Logger : func to Generate Logger in Console and File
func Logger(errortype string, message string, function string) {
	var filename string
	filename = "app.log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	log.SetReportCaller(true)

	if err != nil {
		fmt.Println(err)
	} else {
		mw := io.MultiWriter(os.Stdout, f)
		log.SetOutput(mw)
	}

	if errortype == "info" {
		log.Info(function + " : " + message)
	} else if errortype == "warning" {
		log.Warning(function + " : " + message)
	} else if errortype == "error" {
		log.Error(function + " : " + message)
	}
}
