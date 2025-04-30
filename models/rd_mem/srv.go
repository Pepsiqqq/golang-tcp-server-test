package rd_mem

import "main/models"

type SRV_RD_MEM struct {
	models.Base
	StartAddr uint32 `xml:"StartAddr"`
	DataLen   uint16 `xml:"DataLen"`
}
