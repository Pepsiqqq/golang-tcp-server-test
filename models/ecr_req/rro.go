package ecr_req

import "main/models"

type RRO_ECR_REQ struct {
	models.Base
	TblVer uint32 `xml:"TblVer"`
	Data   []byte `xml:"Data"`
}
