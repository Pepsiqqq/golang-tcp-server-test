package params

import "main/models"

type RRO_PARAMS struct {
	models.Base
	Data []byte `xml:"Data"`
}
