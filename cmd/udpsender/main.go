package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const portNumber = "localhost:42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", portNumber)
	if err != nil {
		fmt.Println("Error resolving udp address: ", err)
	}

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Encountered error reading %s", err)

		}

		_, err = udpConn.Write([]byte(text))
		if err != nil {
			fmt.Printf("Encountered error writing %s", err)
		}

	}

}
