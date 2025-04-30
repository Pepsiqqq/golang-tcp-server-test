package ctrl_req

import (
	"encoding/xml"
	"main/models"
	"strconv"

	"github.com/pkg/errors"
)

type RRO struct {
	models.Base
	ProtType uint16 `xml:"ProtType"`
	DatLen   uint32 `xml:"DatLen"`
	Data     []byte `xml:"Data"`
}

// CreateTestPacket will create test model with test values and marshal it to xml
func (r RRO) CreateTestPacket() ([]byte, error) {
	bs := []byte(strconv.Itoa(200)) // test value
	base, err := r.Base.New(models.MID_RRO_CTRL_REQ)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create base model")
	}

	rro := RRO{
		Base:     base,
		ProtType: 1,
		DatLen:   1,
		Data:     bs,
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
