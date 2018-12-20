package main

import (
	"Alison/createFile"
	"Alison/report"
	"Alison/runServer"
)

func main() {
	createFile.CreateFile()
	go report.PrintReport()
	runServer.RunServer()
}
