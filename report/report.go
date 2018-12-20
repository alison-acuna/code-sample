package report

import (
	"Alison/runServer"
	"fmt"
	"time"
)

func PrintReport() {
	// Every 10 seconds, the Application must print a report to standard output:
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			// The total number of unique numbers received for this run of the Application.
			count := len(runServer.NumbersSlice)
			// The difference since the last report of the count of new unique numbers that have been received.
			uniques := len(runServer.NumbersSlice) - count
			// The difference since the last report of the count of new duplicate numbers that have been received.
			report := fmt.Sprintf("Received %v unique numbers, %v duplicates. Unique total: %v", uniques, runServer.DupCount, count)
			fmt.Println(report)
		}
	}
}
