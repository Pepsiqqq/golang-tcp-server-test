package models

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/go-faster/errors"
)

type RRO_СOM_INI struct {
	Base
	ID_DEV uint32 `xml:"ID_DEV"`
	KSize  uint16 `xml:"KSize"`
	G      []byte `xml:"G"`
	P      []byte `xml:"P"`
	A      []byte `xml:"A"`
	MSize  uint16 `xml:"MSize"`
	MAC    []byte `xml:"MAC"`
	MACKey []byte `xml:"MACKey"`
}

func (r RRO_СOM_INI) Validate() error {
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

func (r RRO_СOM_INI) ValidateKeys() error {
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

func (r RRO_СOM_INI) ValidateMAC() error {
	mSize := (r.MSize + 7) / 8
	if cap(r.MAC) != int(mSize) {
		return errors.New("Error validating rsa size, mSize: " + fmt.Sprint(mSize) + " != " + "Mac size: " + fmt.Sprint(cap(r.MAC)))
	}
	if cap(r.MACKey) != int(mSize) {
		return errors.New("Error validating rsa size, mSize: " + fmt.Sprint(mSize) + " != " + "Mackey size: " + fmt.Sprint(cap(r.MACKey)))
	}
	return nil
}

func (r RRO_СOM_INI) CreateTestPacket() []byte {
	bs := []byte(strconv.Itoa(200))
	base := Base{
		V:       1,
		ProtVer: 2,
		Length:  2,
		SeqNum:  2,
		MID:     0x0001,
		TOut:    2,
		Session: 2,
		Flags:   2,
		ZPad:    0,
		CRC:     1,
	}
	rro := RRO_СOM_INI{
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
	header := `<?xml version="1.0" encoding="windows-1251" ?>` + "\n"
	bytearray, err := xml.MarshalIndent(rro, "", "   ")
	bytearray = []byte(header + string(bytearray))
	if err != nil {
		panic(err)
	}
	return bytearray
}
