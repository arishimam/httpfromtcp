package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("===========================================")

	currentLine := ""
	for {
		buffer := make([]byte, 8, 8)
		n, err := file.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err)
			break

		}
		str := string(buffer[:n])
		parts := strings.Split(str, "\n")

		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s%s\n", currentLine, parts[i])
			currentLine = ""

		}

		currentLine += parts[len(parts)-1]
	}

}
