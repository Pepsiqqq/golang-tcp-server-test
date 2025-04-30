package com_ini

import (
	"encoding/xml"
	"fmt"
	"main/models"
	"strconv"

	"github.com/go-faster/errors"
)

type RRO struct {
	models.Base
	ID_DEV uint32 `xml:"ID_DEV"`
	KSize  uint16 `xml:"KSize"`
	G      []byte `xml:"G"`
	P      []byte `xml:"P"`
	A      []byte `xml:"A"`
	MSize  uint16 `xml:"MSize"`
	MAC    []byte `xml:"MAC"`
	MACKey []byte `xml:"MACKey"`
}

// Valdiate will validate models values
func (r RRO) Validate() error {
	err := r.ValidateKeys()
	if err != nil {
		return errors.Wrap(err, "Failed to validate keys")
	}
	err = r.ValidateMAC()
	if err != nil {
		return errors.Wrap(err, "Failed to validate mac")
	}

	return nil
}

func (r RRO) ValidateKeys() error {
	keySize := (r.KSize + 7) / 8
	if cap(r.G) != int(keySize) {
		return errors.New("Error validating key size, keySize: " + fmt.Sprint(keySize) + " != " + "g size: " + fmt.Sprint(cap(r.G)))
	}
	if cap(r.P) != int(keySize) {
		return errors.New("Error validating key size, keySize: " + fmt.Sprint(keySize) + " != " + "p size: " + fmt.Sprint(cap(r.P)))
	}
	if cap(r.A) != int(keySize) {
		return errors.New("Error  validating key size, keySize: " + fmt.Sprint(keySize) + " != " + "a size: " + fmt.Sprint(cap(r.A)))
	}
	return nil
}

func (r RRO) ValidateMAC() error {
	mSize := (r.MSize + 7) / 8
	if cap(r.MAC) != int(mSize) {
		return errors.New("Error validating rsa size, mSize: " + fmt.Sprint(mSize) + " != " + "Mac size: " + fmt.Sprint(cap(r.MAC)))
	}
	if cap(r.MACKey) != int(mSize) {
		return errors.New("Error validating rsa size, mSize: " + fmt.Sprint(mSize) + " != " + "Mackey size: " + fmt.Sprint(cap(r.MACKey)))
	}
	return nil
}

// CreateTestPacket will create test model with test values and marshal it to xml
func (r RRO) CreateTestPacket() ([]byte, error) {
	bs := []byte(strconv.Itoa(200)) // test value
	base, err := r.Base.New(models.MID_RRO_COM_INI)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create base model")
	}

	rro := RRO{
		Base:   base,
		ID_DEV: 2,
		KSize:  64,
		G:      bs,
		P:      bs,
		A:      bs,
		MSize:  64,
		MAC:    bs,
		MACKey: bs,
	}

	bytearray, err := xml.MarshalIndent(rro, "", "   ")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshal")
	}
	bytearray = []byte(base.GetHeader() + string(bytearray))

	return bytearray, nil
}
