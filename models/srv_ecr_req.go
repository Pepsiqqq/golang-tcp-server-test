package models

type SRV_ECR_REQ struct {
	Base
	TblVer uint32 `xml:"TblVer"`
	Mode   uint8  `xml:"Mode"`
	Data   []byte `xml:"Data"`
}
