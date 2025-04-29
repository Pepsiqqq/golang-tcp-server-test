package models

type RRO_RD_MEM struct {
	Base
	DataLen uint16 `xml:"DataLen"`
	Data    []byte `xml:"Data"`
}
