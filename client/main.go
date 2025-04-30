package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"main/models"
	"main/models/com_ini"
	"main/models/status"
	"net"
	"strconv"

	"golang.org/x/net/html/charset"
)

func main() {
	startClient()
}

// Creates connection to tcp server, sends test packet of choice and listens to responce.
func startClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Read user input for testing model
	fmt.Println("What models do you want to test? Variants:")
	fmt.Println("1.com_ini")
	fmt.Println("2.dta_snd")
	fmt.Println("3.dta_req")
	fmt.Println("4.ctrl_req")
	fmt.Println("5.ecr_req")
	fmt.Println("6.prg_mem")
	fmt.Println("7.rd_mem")
	fmt.Println("8.status")
	fmt.Println("9.params")
	
	var input string
	fmt.Scan(&input)
	modelNumber, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error when convering to string:", err)
		return
	}

	// Creating map[int]struct to use with input
	type mo struct {
		RRO models.Model
		SRV models.Model
	}
	modelsMap := map[int]mo{
		1: {&com_ini.RRO{}, &com_ini.SRV{}},
		// 2: {&com_ini.RRO{}, &com_ini.SRV{}},
		// 3: {&com_ini.RRO{}, &com_ini.SRV{}},
		// 4: {&com_ini.RRO{}, &com_ini.SRV{}},
		// 5: {&com_ini.RRO{}, &com_ini.SRV{}},
		// 6: {&com_ini.RRO{}, &com_ini.SRV{}},
		// 7: {&com_ini.RRO{}, &com_ini.SRV{}},
		8: {&status.RRO{}, &status.SRV{}},
		//9: {&com_ini.RRO{}, &com_ini.SRV{}},
	}

	// Creating model struct based on user input
	model, ok := modelsMap[modelNumber]
	if !ok {
		fmt.Println("Invalid model selected")
		return
	}

	// Sending RRO model to server and listening back for server model
	sendData(conn, model.RRO)
	listenResponce(conn, model.SRV)
}

// sendData will create test packet of selected model and send it to server
func sendData(conn net.Conn, m models.Model) {
	data, err := m.CreateTestPacket()
	if err != nil {
		fmt.Println("Error creating test packet:", err)
		return
	}
	fmt.Println("Sending data to server")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error sendin data to server:", err)
		return
	}
	fmt.Println("Data sended succesfully")
}

// listenResponce will wait for responce from server, unmarshal it into srv model of choice and validate it
func listenResponce(conn net.Conn, m models.Model) {
	fmt.Println("Listening to responce")
	//make buffer for server responce
	buffer := make([]byte, 4096)
	_, err := conn.Read(buffer)
	fmt.Println("Readed responce")
	if err != nil && err != io.EOF {
		fmt.Println("Error1:", err)
		return
	}
	fmt.Print("Got this xml from server:")
	fmt.Println(string(buffer))
	fmt.Println("Unmarshaling responce")
	// create new xml decoder set character set to windows-1251 and decode xml into srv model
	decoder := xml.NewDecoder(bytes.NewReader(buffer))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&m)
	if err != nil {
		fmt.Println("Error2:", err)
		return
	}
	fmt.Println("Done")
	fmt.Println("Validating")
	err = m.Validate()
	if err != nil {
		fmt.Println("Error while validating:", err)
		return
	}
	fmt.Println("Done")
}
