package models

import (
	"fmt"

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
