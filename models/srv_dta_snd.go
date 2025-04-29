package models

import (
	"encoding/xml"
)

type SRV_DTA_SND struct {
	Base
	Result uint32 `xml:"Result"`
}

func (r SRV_DTA_SND) CreateTestPacket() []byte {
	//bs := []byte(strconv.Itoa(200))
	base := Base{
		V:       1,
		ProtVer: 2,
		Length:  2,
		SeqNum:  2,
		MID:     0x0006,
		TOut:    2,
		Session: 2,
		Flags:   2,
		ZPad:    0,
		CRC:     1,
	}
	rro := SRV_DTA_SND{
		Base:   base,
		Result: 1,
	}
	header := `<?xml version="1.0" encoding="windows-1251" ?>` + "\n"
	bytearray, err := xml.MarshalIndent(rro, "", "   ")
	bytearray = []byte(header + string(bytearray))
	if err != nil {
		panic(err)
	}
	return bytearray
}
