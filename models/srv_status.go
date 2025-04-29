package models

import (
	"encoding/xml"

)

type SRV_STATUS struct {
	XMLName xml.Name `xml:"RS"`
	V       int      `xml:"V,attr"`
	Base
	Data []byte `xml:"Data"`
	ZPad uint32 `xml:"ZPad"`
	CRC  uint32 `xml:"CRC"`
}

func (r SRV_STATUS) Validate() error{
	return nil
}