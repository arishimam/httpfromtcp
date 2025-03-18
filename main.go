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
				currentLine += string(data[:i])
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
				if i+1 < len(data) {
					data = data[i+1:]
				} else {
					data = []byte{}
				}
				break
			}

		}
		currentLine += string(data)
	}

}
