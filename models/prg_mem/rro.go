package prg_mem

import "main/models"

type RRO_PRG_MEM struct {
	models.Base
	DataLen uint16 `xml:"DataLen"`
}
