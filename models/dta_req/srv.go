package dta_req

import "main/models"

type SRV struct {
	models.Base
	PackType uint8  `xml:"PackType"`
	PackNum  uint8  `xml:"PackNum"`
	Error    uint16 `xml:"Error"`
	CTime    uint64 `xml:"CTime"`
	STime    uint64 `xml:"STime"`
	FTime    uint64 `xml:"FTime"`
	DataLen  uint32 `xml:"DataLen"`
	Service  uint8  `xml:"Service"`
	ShiftNum uint16 `xml:"ShiftNum"`
	DocNum   uint32 `xml:"DocNum"`
	Upload   uint32 `xml:"Upload"`
	Data     []byte `xml:"Data"`
}
