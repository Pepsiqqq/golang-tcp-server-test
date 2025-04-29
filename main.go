package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net"

	"golang.org/x/net/html/charset"
	"main/models"
)

func main() {

	startServer()
}

func startServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 4096)
	_, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
		return
	}
	base := models.Base{}
	fmt.Println(string(buffer))
	decoder := xml.NewDecoder(bytes.NewReader(buffer))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&base)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = ValidateModel(base)
	if err != nil {
		fmt.Println("Validation failed", err)
		return
	}
	var model struct{}
	switch base.MID{
	case 0x0001:
		err = decoder.Decode(&model)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}	
		test := models.CreateTestPacket()
		_, err = conn.Write(test)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
}
