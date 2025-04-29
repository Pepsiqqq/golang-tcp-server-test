package models

type SRV_CTRL_REQ struct {
	Base
	DatLen   uint32 `xml:"DatLen"`
	Data     []byte `xml:"Data"`
}
