package main

import (
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")

	if err != nil {
		log.Println(err)
		return
	} else {
		defer sysLog.Close()
		log.SetOutput(sysLog)
		log.Println("I am thinking syslog is work")
	}
}
