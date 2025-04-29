package models

import "encoding/xml"

type SRV_DTA_SND struct {
	XMLName xml.Name `xml:"RS"`
	V       int      `xml:"V,attr"`
	Base
	ID_DEV  uint32   `xml:"ID_DEV"`
	B       uint32   `xml:"G"`
	MSize   uint16   `xml:"MSize"`
	MAC     uint32   `xml:"MAC"`
	MACKey  uint32   `xml:"MACKey"`
	ZPad    uint32   `xml:"ZPad"`
	CRC     uint32   `xml:"CRC"`
}