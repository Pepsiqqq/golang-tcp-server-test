package status

import (
	"encoding/xml"
	"main/models"
	"strconv"

	"github.com/pkg/errors"
)

type RRO struct {
	models.Base
	Data []byte `xml:"Data"`
}

// CreateTestPacket will create test model with test values and marshal it to xml
func (r RRO) CreateTestPacket() ([]byte, error) {
	bs := []byte(strconv.Itoa(200)) // test value
	base, err := r.Base.New(models.MID_RRO_STATUS)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create base model")
	}

	rro := RRO{
		Base: base,
		Data: bs,
	}

	bytearray, err := xml.MarshalIndent(rro, "", "   ")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshal")
	}
	bytearray = []byte(base.GetHeader() + string(bytearray))

	return bytearray, nil
}

func (r RRO) Validate() error {
	//TODO - implement
	return nil
}
