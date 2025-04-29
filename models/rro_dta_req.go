package models

type RRO_DTA_REQ struct {
	Base
	SearchBy uint8  `xml:"SearchBy"`
	InfoFrom uint64 `xml:"InfoFrom"`
	InfoTo   uint64 `xml:"InfoTo"`
}
