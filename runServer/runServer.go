package runServer

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var NumbersSlice []string
var DupCount int

func manageInput(c net.Conn, k chan string, l net.Listener) {

	filename := "./numbers.log"

	for {
		select {
		default:
			// UI for number input
			reader := bufio.NewReader(c)
			c.Write([]byte(string("Enter text: ")))
			Text, _ := reader.ReadString('\n')

			dup := false

			// Checks to ensure input is formatted correctly
			temp := strings.TrimSpace(Text)
			_, err1 := strconv.Atoi(temp)
			textCounted := len(temp)

			if temp == "terminate" {
				k <- "kill"
				c.Close()
				l.Close()
			} else if err1 != nil {
				openConnection(k, l)
				c.Close()
			} else if textCounted != 9 {
				openConnection(k, l)
				c.Close()
			}

			// Checks to ensure input is not a duplicate
			for _, elem2 := range NumbersSlice {
				if Text == elem2 {
					dup = true
				}
			}

			if dup == false {
				NumbersSlice = append(NumbersSlice, Text)
				f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				if _, err = f.WriteString(Text); err != nil {
					panic(err)
				}
			} else {
				DupCount++
			}

		case <-k:
			// In case terminate channel includes message, close channel
			c.Close()
		}
	}
	c.Close()
}

func RunServer() {
	// Define channel to communicate terminate sequence
	killSwitch := make(chan string)

	// Open server on port 4000
	PORT := ":4000"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	DupCount = 0

	// Define channel to limit number of connections
	maxGoroutines := 5
	guard := make(chan struct{}, maxGoroutines)

	// Opens a connection and calls function to manage input
	for {
		guard <- struct{}{}
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go manageInput(c, killSwitch, l)
	}
}

func openConnection(k chan string, l net.Listener) {

	maxGoroutines := 1
	guard := make(chan struct{}, maxGoroutines)

	// Opens a connection and calls function to manage input
	for {
		guard <- struct{}{}
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go manageInput(c, k, l)
	}
}
