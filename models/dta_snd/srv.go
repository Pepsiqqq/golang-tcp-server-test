package dta_snd

import (
	"encoding/xml"
	"main/models"

	"github.com/pkg/errors"
)

type SRV struct {
	models.Base
	Result uint32 `xml:"Result"`
}

func (s SRV) CreateTestPacket() ([]byte, error) {
	//bs := []byte(strconv.Itoa(200)) // test value
	base, err := s.Base.New(models.MID_SRV_DTA_SND)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create base model")
	}
	srv := SRV{
		Base:   base,
		Result: 1,
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
