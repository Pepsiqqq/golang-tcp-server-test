package rd_mem

import "main/models"

type RRO_RD_MEM struct {
	models.Base
	DataLen uint16 `xml:"DataLen"`
	Data    []byte `xml:"Data"`
}
