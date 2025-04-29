package models

import (
	"encoding/xml"
	"strconv"
)

type SRV_СOM_INI struct {
	XMLName xml.Name `xml:"RS"`
	V       int      `xml:"V,attr"`
	Base
	ID_DEV uint32 `xml:"ID_DEV"`
	B      []byte `xml:"B"`
	MSize  uint16 `xml:"MSize"`
	MAC    []byte `xml:"MAC"`
	MACKey []byte `xml:"MACKey"`
	ZPad   uint32 `xml:"ZPad"`
	CRC    uint32 `xml:"CRC"`
}

func CreateTestPacket() []byte {
	bs := []byte(strconv.Itoa(200))
	base := Base{
		ProtVer: 1,
		Length:  1,
		SeqNum:  1,
		MID:     0x0002,
		TOut:    1,
		Session: 1,
		Flags:   1,
	}
	s := SRV_СOM_INI{
		V:      1,
		Base:   base,
		ID_DEV: 1,
		B:      bs,
		MSize:  64,
		MAC:    bs,
		MACKey: bs,
		ZPad:   1,
		CRC:    1,
	}
	header := `<?xml version="1.0" encoding="windows-1251" ?>` + "\n"
	bytearray, err := xml.MarshalIndent(s, "", "   ")
	bytearray = []byte(header + string(bytearray))
	if err != nil {
		panic(err)
	}
	return bytearray
}
