package dta_snd

import (
	"encoding/xml"
	"main/models"
	"strconv"

	"github.com/go-faster/errors"
)

type RRO struct {
	models.Base
	PackType uint8  `xml:"PackType"`
	PackNum  uint8  `xml:"PackNum"`
	Error    uint16 `xml:"Error"`
	CTime    []byte `xml:"CTime"`
	STime    uint64 `xml:"STime"`
	FTime    uint64 `xml:"FTime"`
	DataLen  uint32 `xml:"DataLen"`
	Service  uint8  `xml:"Service"`
	ShiftNum uint16 `xml:"ShiftNum"`
	DocNum   uint32 `xml:"DocNum"`
	Upload   uint32 `xml:"Upload"`
	Data     []byte `xml:"Data"`
}

func (r RRO) Validate() error {
	mid := map[int]struct{}{
		0x11: {},
		0x12: {},
		0x13: {},
	}
	_, ok := mid[int(r.PackType)]
	if !ok {
		return errors.New("invalid PackType")
	}

	r.validateDates()
	return nil
}

func (r RRO) validateDates() {
	// TODO - implement
	//loc := time.Local
	// day := r.CTime[0:1]
	// month := r.CTime[1:2]
	// year := r.CTime[2:3]
	// hour := r.CTime[3:4]
	// min := r.CTime[4:5]
	// fmt.Print(string(day), string(month), string(year), string(hour), string(min))
	//return time.Date(year, month, day, hour, min, nil, nil, loc)
}

func (r RRO) CreateTestPacket() ([]byte, error) {
	bs := []byte(strconv.Itoa(200)) // test value
	base, err := r.Base.New(models.MID_RRO_COM_INI)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create base model")
	}
	rro := RRO{
		Base:     base,
		PackType: 0x11,
		PackNum:  1,
		Error:    1,
		CTime:    bs,
		STime:    1,
		FTime:    1,
		DataLen:  1,
		Service:  1,
		ShiftNum: 1,
		DocNum:   1,
		Upload:   1,
		Data:     bs,
	}
	bytearray, err := xml.MarshalIndent(rro, "", "   ")
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshal")
	}
	bytearray = []byte(base.GetHeader() + string(bytearray))

	return bytearray, nil
}
