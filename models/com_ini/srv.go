package com_ini

import (
	"encoding/xml"
	"main/models"
	"strconv"

	"github.com/go-faster/errors"
)

type SRV struct {
	models.Base
	ID_DEV uint32 `xml:"ID_DEV"`
	B      []byte `xml:"B"`
	MSize  uint16 `xml:"MSize"`
	MAC    []byte `xml:"MAC"`
	MACKey []byte `xml:"MACKey"`
}

// CreateTestPacket will create test model with test values and marshal it to xml
func (s SRV) CreateTestPacket() ([]byte, error) {
	bs := []byte(strconv.Itoa(200)) // test value
	base, err := s.Base.New(models.MID_SRV_COM_INI)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create base model")
	}

	srv := SRV{
		Base:   base,
		ID_DEV: 1,
		B:      bs,
		MSize:  64,
		MAC:    bs,
		MACKey: bs,
	}
	bytearray, err := xml.MarshalIndent(srv, "", "   ")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshal")
	}
	bytearray = []byte(base.GetHeader() + string(bytearray))

	return bytearray, nil
}

func (r SRV) Validate() error {
	//TODO
	return nil
}
