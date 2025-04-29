package models

import (
	"encoding/xml"
	"strconv"
)

type SRV_СOM_INI struct {
	Base
	ID_DEV uint32 `xml:"ID_DEV"`
	B      []byte `xml:"B"`
	MSize  uint16 `xml:"MSize"`
	MAC    []byte `xml:"MAC"`
	MACKey []byte `xml:"MACKey"`
}

func (s SRV_СOM_INI) CreateTestPacket() []byte {
	bs := []byte(strconv.Itoa(200))
	base := Base{
		V:      1,
		ProtVer: 1,
		Length:  1,
		SeqNum:  1,
		MID:     0x0002,
		TOut:    1,
		Session: 1,
		Flags:   1,
		ZPad:   1,
		CRC:    1,
	}
	y := SRV_СOM_INI{
		Base:   base,
		ID_DEV: 1,
		B:      bs,
		MSize:  64,
		MAC:    bs,
		MACKey: bs,
	}
	header := `<?xml version="1.0" encoding="windows-1251" ?>` + "\n"
	bytearray, err := xml.MarshalIndent(y, "", "   ")
	bytearray = []byte(header + string(bytearray))
	if err != nil {
		panic(err)
	}
	return bytearray
}

func (r SRV_СOM_INI) Validate() error{
	return nil
}