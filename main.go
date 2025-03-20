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
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Error opening: %v", err)
	}

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("================================")

	linesChan := getLinesChannel(f)

	for line := range linesChan {
		fmt.Printf("read: %s\n", line)
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {

	ch := make(chan string)

	go func() {
		defer close(ch)
		defer f.Close()

		currentLine := ""
		for {
			buffer := make([]byte, 8, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if currentLine != "" {
					ch <- currentLine
				}

				if errors.Is(err, io.EOF) {
					break
				}

				fmt.Printf("Error reading the file: %v", err)
				break
			}

			str := string(buffer[:n])
			parts := strings.Split(str, "\n")

			for i := 0; i < len(parts)-1; i++ {
				currentLine += parts[i]
				ch <- currentLine
				currentLine = ""
			}

			currentLine += parts[len(parts)-1]

		}

	}()

	return ch

}
