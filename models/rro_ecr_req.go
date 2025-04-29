package models


type RRO_ECR_REQ struct {
	Base
	TblVer uint32 `xml:"TblVer"`
	Data   []byte `xml:"Data"`
}
