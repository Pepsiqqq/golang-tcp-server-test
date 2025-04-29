package models

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
)

type RRO_DTA_SND struct {
	Base
	PackType uint8  `xml:"PackType"`
	PackNum  uint8  `xml:"PackNum"`
	Error    uint16 `xml:"Error"`
	CTime    []byte `xml:"CTime"`
	STime    uint64 `xml:"STime"`
	FTime    uint64 `xml:"FTime"`
	DataLen  uint32 `xml:"DataLen"`
	Service  uint8  `xml:"Service"`
	ShiftNum uint16 `xml:"ShiftNum"`
	DocNum   uint32 `xml:"DocNum"`
	Upload   uint32 `xml:"Upload"`
	Data     []byte `xml:"Data"`
}

func (r RRO_DTA_SND) Validate() error {
	switch r.PackType {
	case 0x11:
	case 0x12:
	case 0x13:
	default:
		return errors.New("PackType error")
	}
	r.parseDate()
	return nil
}

func (r RRO_DTA_SND) parseDate() {
	//loc := time.Local
	day := r.CTime[0:1]
	month := r.CTime[1:2]
	year := r.CTime[2:3]
	hour := r.CTime[3:4]
	min := r.CTime[4:5]
	fmt.Print(string(day), string(month), string(year), string(hour), string(min))
	//return time.Date(year, month, day, hour, min, nil, nil, loc)
}

func (r RRO_DTA_SND) CreateTestPacket() []byte {
	bs := []byte(strconv.Itoa(200))
	bs2 := make([]byte, 5)
	i := 24
	bs2[0] = byte(i)
	bs2[1] = byte(i)
	bs2[2] = byte(i)
	bs2[3] = byte(i)
	bs2[4] = byte(i)
	base := Base{
		V:       1,
		ProtVer: 2,
		Length:  2,
		SeqNum:  2,
		MID:     0x0005,
		TOut:    2,
		Session: 2,
		Flags:   2,
		ZPad:    0,
		CRC:     1,
	}
	rro := RRO_DTA_SND{
		Base:     base,
		PackType: 0x11,
		PackNum:  1,
		Error:    1,
		CTime:    bs2,
		STime:    1,
		FTime:    1,
		DataLen:  1,
		Service:  1,
		ShiftNum: 1,
		DocNum:   1,
		Upload:   1,
		Data:     bs,
	}
	header := `<?xml version="1.0" encoding="windows-1251" ?>` + "\n"
	bytearray, err := xml.MarshalIndent(rro, "", "   ")
	bytearray = []byte(header + string(bytearray))
	if err != nil {
		panic(err)
	}
	return bytearray
}
