package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net"

	"main/models"

	"golang.org/x/net/html/charset"
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
	basebuffer := make([]byte, 4096)
	modelbuffer := make([]byte, 4096)
	_, err := conn.Read(basebuffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
		return
	}
	modelbuffer = basebuffer
	base := models.Base{}
	fmt.Println(string(basebuffer))
	decoder := xml.NewDecoder(bytes.NewReader(basebuffer))
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
	switch base.MID {
	case 0x0001:
		fmt.Println(string(modelbuffer))
		modelRRO := models.RRO_СOM_INI{}
		decoder := xml.NewDecoder(bytes.NewReader(modelbuffer))
		decoder.CharsetReader = charset.NewReaderLabel
		err = decoder.Decode(&modelRRO)
		if err != nil {
			fmt.Println("Decode error:", err)
			return
		}
		err = modelRRO.Validate()
		if err != nil {
			fmt.Println("Validating error:", err)
			return
		}
		modelSRV := models.SRV_СOM_INI{}
		test := modelSRV.CreateTestPacket()
		_, err = conn.Write(test)
		if err != nil {
			fmt.Println("Writing to client error:", err)
			return
		}
	case 0x00011:
		fmt.Println(string(modelbuffer))
		modelRRO := models.RRO_STATUS{}
		decoder := xml.NewDecoder(bytes.NewReader(modelbuffer))
		decoder.CharsetReader = charset.NewReaderLabel
		err = decoder.Decode(&modelRRO)
		if err != nil {
			fmt.Println("Decode error:", err)
			return
		}
		err = modelRRO.Validate()
		if err != nil {
			fmt.Println("Validating error:", err)
			return
		}
		modelSRV := models.SRV_STATUS{}
		test := modelSRV.CreateTestPacket()
		_, err = conn.Write(test)
		if err != nil {
			fmt.Println("Writing to client error:", err)
			return
		}
	}
}
