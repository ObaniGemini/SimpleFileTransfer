package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	InitGUI()

	// Connect to the server
	reader := bufio.NewReader(os.Stdin)

	// Get the receiver's IP address
	fmt.Print("Enter receiver's IP address: ")
	receiverIP, _ := reader.ReadString('\n')
	receiverIP = strings.TrimSpace(receiverIP) // Remove any trailing newline

	// Get the receiver's port
	fmt.Print("Enter receiver's Port: ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(receiverIP) // Remove any trailing newline

	// Get the file path
	fmt.Print("Enter path to the file: ")
	filepath, _ := reader.ReadString('\n')
	filepath = strings.TrimSpace(filepath) // Remove any trailing newline

	// Connect to the server
	conn, err := net.Dial("tcp", receiverIP+":"+port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Open the file to be sent
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Copy the file to the connection
	_, err = io.Copy(conn, file)
	if err != nil {
		panic(err)
	}

	fmt.Println("This program does nothing for now")
}
