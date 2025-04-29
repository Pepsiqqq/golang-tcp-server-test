package models

import (
	"encoding/xml"
	"strconv"
)

type SRV_STATUS struct {
	Base
	Data []byte `xml:"Data"`
}

func (r SRV_STATUS) Validate() error {
	
	return nil
}

func (s SRV_STATUS) CreateTestPacket() []byte {
	bs := []byte(strconv.Itoa(200))
	base := Base{
		V:       1,
		ProtVer: 1,
		Length:  1,
		SeqNum:  1,
		MID:     0x0012,
		TOut:    1,
		Session: 1,
		Flags:   1,
		ZPad:    1,
		CRC:     1,
	}
	y := SRV_STATUS{
		Base: base,
		Data: bs,
	}
	header := `<?xml version="1.0" encoding="windows-1251" ?>` + "\n"
	bytearray, err := xml.MarshalIndent(y, "", "   ")
	bytearray = []byte(header + string(bytearray))
	if err != nil {
		panic(err)
	}
	return bytearray
}
