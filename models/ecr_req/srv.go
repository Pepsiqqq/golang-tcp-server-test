package ecr_req

import "main/models"

type SRV_ECR_REQ struct {
	models.Base
	TblVer uint32 `xml:"TblVer"`
	Mode   uint8  `xml:"Mode"`
	Data   []byte `xml:"Data"`
}
