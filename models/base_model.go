package models

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type Base struct {
	XMLName xml.Name `xml:"RS"`
	V       int      `xml:"V,attr"`
	ProtVer uint32 `xml:"ProtVer"`
	Length  uint32 `xml:"Length"`
	SeqNum  uint8  `xml:"SeqNum"`
	MID     uint16 `xml:"MID"`
	TOut    uint16 `xml:"TOut"`
	Session uint16 `xml:"Session"`
	Flags   uint16 `xml:"Flags"`
	ZPad    uint32 `xml:"ZPad"`
	CRC     uint32 `xml:"CRC"`
}

func (b Base) Validate() error {
	fmt.Print("Validating Base")
	if b.MID != 0x0001 && b.MID != 0x0002 {
		if b.MID < 0x0005 && b.MID > 0x0014 {
			return errors.New("invalid MID")
		}
	}
	switch b.Flags {
	case 0x0001:
	case 0x0002:
	case 0x0004:
	case 0x0008:
	case 0x0010:
	case 0x0020:
	case 0x0040:
	case 0x0080:
	case 0x0100:
	case 0x8000:
	default:
		return errors.New("invalid Flags")
	}

	if b.ZPad != 0 {
		return errors.New("zpad not null")
	}
	return nil
}
