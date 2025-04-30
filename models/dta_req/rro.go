package dta_req

import (
	"encoding/xml"
	"main/models"

	"github.com/pkg/errors"
)

type RRO struct {
	models.Base
	SearchBy uint8  `xml:"SearchBy"`
	InfoFrom uint64 `xml:"InfoFrom"`
	InfoTo   uint64 `xml:"InfoTo"`
}

// CreateTestPacket will create test model with test values and marshal it to xml
func (r RRO) CreateTestPacket() ([]byte, error) {
	//bs := []byte(strconv.Itoa(200)) // test value
	base, err := r.Base.New(models.MID_RRO_DTA_REQ)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create base model")
	}
	rro := RRO{
		Base:     base,
		SearchBy: 1,
		InfoFrom: 1,
		InfoTo:   1,
	}

	bytearray, err := xml.MarshalIndent(rro, "", "   ")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshal")
	}
	bytearray = []byte(base.GetHeader() + string(bytearray))

	return bytearray, nil
}

// Valdiate will validate models values
func (r RRO) Validate() error {
	//TODO - create validation
	return nil
}
