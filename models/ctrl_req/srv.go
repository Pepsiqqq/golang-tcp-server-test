package ctrl_req

import (
	"encoding/xml"
	"main/models"
	"strconv"

	"github.com/pkg/errors"
)

type SRV struct {
	models.Base
	DatLen uint32 `xml:"DatLen"`
	Data   []byte `xml:"Data"`
}

// CreateTestPacket will create test model with test values and marshal it to xml
func (s SRV) CreateTestPacket() ([]byte, error) {
	bs := []byte(strconv.Itoa(200)) // test value
	base, err := s.Base.New(models.MID_SRV_CTRL_REQ)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create base model")
	}
	srv := SRV{
		Base:   base,
		DatLen: 1,
		Data:   bs,
	}
	bytearray, err := xml.MarshalIndent(srv, "", "   ")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshal")
	}
	bytearray = []byte(base.GetHeader() + string(bytearray))

	return bytearray, nil
}

func (r SRV) Validate() error {
	//TODO - create validation for server packets
	return nil
}
