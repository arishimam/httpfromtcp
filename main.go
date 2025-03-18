package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	currentLine := ""
	for {
		data := make([]byte, 8)
		_, err := file.Read(data)
		if err != nil {
			break
		}

		for i := range data {

			if data[i] == '\n' {
				// split
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
				continue

			}
			currentLine += string(data[i])
		}
	}

	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}

}
