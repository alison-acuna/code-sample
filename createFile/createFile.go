package createFile

import (
	"log"
	"os"
)

func CreateFile() {
	f, err := os.Create("numbers.log")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer f.Close()
}
