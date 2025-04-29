package models

type RRO_STATUS struct {
	Base
	Data []byte `xml:"Data"`
}

func (r RRO_STATUS) Validate() error{
	return nil
}