package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net"

	"main/models"
	"main/models/com_ini"

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
	// create two identical buffers, read from one to get MID and set model, read from second to get model
	basebuffer := make([]byte, 4096)
	modelbuffer := basebuffer
	_, err := conn.Read(basebuffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
		return
	}
	base := models.Base{}
	fmt.Println("Get this XML:")
	fmt.Println(string(basebuffer))
	decoder := xml.NewDecoder(bytes.NewReader(basebuffer))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&base)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = base.Validate()
	if err != nil {
		fmt.Println("Validation failed", err)
		return
	}
	modelRROMap := map[int]models.Model{
		0x0001: com_ini.RRO{},
		// 0x0005: {},
		// 0x0007: {},
		// 0x0009: {},
		// 0x000b: {},
		// 0x000d: {},
		// 0x0010: {},
		// 0x0012: {},
	}
	modelRRO := modelRROMap[int(base.MID)] // don't need to check ok because already validated base model

	decoder = xml.NewDecoder(bytes.NewReader(modelbuffer))
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
	// prepare and send test reply
	fmt.Println("Send reply")
	modelSRVMap := map[int]models.Model{
		0x0002: com_ini.SRV{},
		// 0x0006: {},
		// 0x0008: {},
		// 0x000a: {},
		// 0x000c: {},
		// 0x000f: {},
		// 0x0011: {},
		// 0x0013: {},
	}
	modelSRV := modelSRVMap[int(base.MID+0x0001)] // don't need to check ok because already validated base model
	test, err := modelSRV.CreateTestPacket()
	if err != nil {
		fmt.Println("Failed to create test packet:", err)
		return
	}
	_, err = conn.Write(test)
	if err != nil {
		fmt.Println("Writing to client error:", err)
		return
	}
}
