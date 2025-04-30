package models

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Base struct {
	XMLName xml.Name `xml:"RS"`
	V       int      `xml:"V,attr"`
	ProtVer uint32   `xml:"ProtVer"`
	Length  uint32   `xml:"Length"`
	SeqNum  uint8    `xml:"SeqNum"`
	MID     uint16   `xml:"MID"`
	TOut    uint16   `xml:"TOut"`
	Session uint16   `xml:"Session"`
	Flags   uint16   `xml:"Flags"`
	ZPad    uint32   `xml:"ZPad"`
	CRC     uint32   `xml:"CRC"`
}

// GetHeader returns default header when marshaling to xml
func (b Base) GetHeader() string {
	return `<?xml version="1.0" encoding="windows-1251" ?>` + "\n"
}

// New creates new model, sets mid and returns it
func (b Base) New(mid uint16) (Base, error) {
	// TODO - transfer all except mid to CreateTestPacket
	base := Base{
		V:       1,
		ProtVer: 2,
		Length:  2,
		SeqNum:  2,
		TOut:    2,
		MID:     mid,
		Session: 2,
		Flags:   0x0001,
		ZPad:    0,
		CRC:     1,
	}
	return base, nil
}

// Valdiate will validate base model values
func (b Base) Validate() error {
	fmt.Println("Validating Base")
	mid := map[int]struct{}{
		MID_RRO_COM_INI:  {},
		MID_SRV_COM_INI:  {},
		MID_RRO_DTA_SND:  {},
		MID_SRV_DTA_SND:  {},
		MID_RRO_DTA_REQ:  {},
		MID_SRV_DTA_REQ:  {},
		MID_RRO_CTRL_REQ: {},
		MID_SRV_CTRL_REQ: {},
		MID_RRO_ECR_REQ:  {},
		MID_SRV_ECR_REQ:  {},
		MID_RRO_PRG_MEM:  {},
		MID_SRV_PRG_MEM:  {},
		MID_RRO_RD_MEM:   {},
		MID_SRV_RD_MEM:   {},
		MID_RRO_STATUS:   {},
		MID_SRV_STATUS:   {},
		MID_RRO_PARAMS:   {},
		MID_SRV_PARAMS:   {},
	}
	_, ok := mid[int(b.MID)]
	if !ok {
		return errors.New("invalid MID")
	}
	flags := map[int]struct{}{
		0x0001: {},
		0x0002: {},
		0x0004: {},
		0x0008: {},
		0x0010: {},
		0x0020: {},
		0x0040: {},
		0x0080: {},
		0x0100: {},
		0x8000: {},
	}
	_, ok = flags[int(b.Flags)]
	if !ok {
		return errors.New("invalid Flags")
	}

	if b.ZPad != 0 {
		return errors.New("zpad not null")
	}
	return nil
}

func (b Base) CreateTestPacket() ([]byte, error) {
	// TODO - implement
	return nil, nil
}
