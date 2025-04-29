package models

type SRV_RD_MEM struct {
	Base
	StartAddr uint32 `xml:"StartAddr"`
	DataLen   uint16 `xml:"DataLen"`
}
