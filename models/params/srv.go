package params

import "main/models"

type SRV_PARAMS struct {
	models.Base
	Data []byte `xml:"Data"`
}
