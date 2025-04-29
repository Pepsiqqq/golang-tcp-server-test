package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"main/models"
	"net"
	"os"

	"golang.org/x/net/html/charset"
)


func main() {
	startClient()
}

func startClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	sendData(conn)
	listenResponce(conn)
}

func sendData(conn net.Conn) {
	modelRRO := models.RRO_СOM_INI{}
	data := modelRRO.CreateTestPacket()
	file, _ := os.Create("test.xml")
	file.Write(data)
	fmt.Println("Sending data")
	_, err := conn.Write(data)
	fmt.Println("Sended data")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func listenResponce(conn net.Conn) {
	fmt.Println("Listening to responce")
	buffer := make([]byte, 4096)
	_, err := conn.Read(buffer)
	fmt.Println("Readed responce")
	if err != nil && err != io.EOF {
		fmt.Println("Error1:", err)
		return
	}
	srv := models.SRV_СOM_INI{}
	fmt.Println(string(buffer))
	fmt.Println("Unmarshaling")
	decoder := xml.NewDecoder(bytes.NewReader(buffer))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&srv)
	if err != nil {
		fmt.Println("Error2:", err)
		return
	}
	fmt.Println("Done")
	if srv.MID == 0x0002 {
		fmt.Printf("All good")
	}
}

