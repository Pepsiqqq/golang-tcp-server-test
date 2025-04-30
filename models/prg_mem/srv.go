package prg_mem

import "main/models"

type SRV_PRG_MEM struct {
	models.Base
	StartAddr uint32 `xml:"StartAddr"`
	DataLen   uint16 `xml:"DataLen"`
	Data      []byte `xml:"Data"`
}
