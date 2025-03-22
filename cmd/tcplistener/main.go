package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	//"os"
	"net"
	"strings"
)

// const inputFilePath = "messages.txt"
const portNumber = ":42069"

func main() {
	// f, err := os.Open(inputFilePath)
	// if err != nil {
	// 	log.Fatalf("Error opening: %v", err)
	// }

	fmt.Printf("Starting tcp listener on %s\n", portNumber)

	listener, err := net.Listen("tcp", portNumber)
	if err != nil {
		log.Fatalf("Error listening: %v", err)
		return
	}
	defer listener.Close()

	fmt.Println("Waiting for connection to be accepted")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection: %v", err)
			return

		}

		fmt.Println("Connection has been accepted")
		fmt.Println("================================")

		linesChan := getLinesChannel(conn)

		for line := range linesChan {
			fmt.Printf("%s\n", line)
		}

	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {

	ch := make(chan string)

	go func() {

		defer close(ch)
		defer fmt.Println("Channel has been closed")
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
