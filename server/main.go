package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net"

	"main/models"
	"main/models/com_ini"
	"main/models/status"

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
		models.MID_RRO_COM_INI: com_ini.RRO{},
		// models.MID_RRO_DTA_SND:  {},
		// models.MID_RRO_DTA_REQ:  {},
		// models.MID_RRO_CTRL_REQ: {},
		// models.MID_RRO_ECR_REQ:  {},
		// models.MID_RRO_PRG_MEM:  {},
		// models.MID_RRO_RD_MEM:   {},
		models.MID_RRO_STATUS: status.RRO{},
		//models.MID_RRO_PARAMS:   {},
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
		models.MID_SRV_COM_INI: com_ini.SRV{},
		// MID_SRV_DTA_SND:  {},
		// MID_SRV_DTA_REQ:  {},
		// MID_SRV_CTRL_REQ: {},
		// MID_SRV_ECR_REQ:  {},
		// MID_SRV_PRG_MEM:  {},
		// MID_SRV_RD_MEM:   {},
		models.MID_SRV_STATUS: status.SRV{},
		// MID_SRV_PARAMS:   {},
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
