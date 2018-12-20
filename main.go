package main

import (
	"newrelic-interviews/Alison/createFile"
	"newrelic-interviews/Alison/report"
	"newrelic-interviews/Alison/runServer"
)

func main() {
	createFile.CreateFile()
	go report.PrintReport()
	runServer.RunServer()
}
