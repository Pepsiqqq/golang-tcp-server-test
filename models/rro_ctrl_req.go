package models

type RRO_CTRL_REQ struct {
	Base
	ProtType uint16 `xml:"ProtType"`
	DatLen   uint32 `xml:"DatLen"`
	Data     []byte `xml:"Data"`
}
