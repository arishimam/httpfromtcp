package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const serverAddress = "localhost:42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		fmt.Println("Error resolving udp address: ", err)
		os.Exit(1)
	}

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Printf("Error dialing UDP %s", err)
		os.Exit(1)
	}
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Encountered error reading %s", err)
			os.Exit(1)
		}

		_, err = udpConn.Write([]byte(text))
		if err != nil {
			fmt.Printf("Encountered error writing %s", err)
			os.Exit(1)
		}
		fmt.Printf("Message sent: %s\n", text)
	}

}
