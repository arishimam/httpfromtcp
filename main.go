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

	ch := getLinesChannel(f)

	for l := range ch {
		fmt.Printf("read: %s\n", l)
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	currentLine := ""

	ch := make(chan string)

	go func() {
		for {
			buffer := make([]byte, 8, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if currentLine != "" {
					fmt.Printf("read: %s\n", currentLine)
					ch <- currentLine

				}

				if errors.Is(err, io.EOF) {
					close(ch)
					break
				}

				fmt.Printf("Error reading the file: %v", err)
				break
			}

			parts := strings.Split(string(buffer[:n]), "\n")

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
