package models

const MID_RRO_COM_INI = 0x0001
const MID_SRV_COM_INI = 0x0002
const MID_RRO_DTA_SND = 0x0005
const MID_SRV_DTA_SND = 0x0006
const MID_RRO_DTA_REQ = 0x0007
const MID_SRV_DTA_REQ = 0x0008
const MID_RRO_CTRL_REQ = 0x0009
const MID_SRV_CTRL_REQ = 0x000a
const MID_RRO_ECR_REQ = 0x000b
const MID_SRV_ECR_REQ = 0x000c
const MID_RRO_PRG_MEM = 0x000d
const MID_SRV_PRG_MEM = 0x000e
const MID_RRO_RD_MEM = 0x000f
const MID_SRV_RD_MEM = 0x0010
const MID_RRO_STATUS = 0x0011
const MID_SRV_STATUS = 0x0012
const MID_RRO_PARAMS = 0x0013
const MID_SRV_PARAMS = 0x0014

type Model interface {
	Validate() error
	CreateTestPacket() ([]byte, error)
}

func ValidateModel(m Model) error {
	return m.Validate()
}
func CreateTestPacketModel(m Model) ([]byte, error) {
	return m.CreateTestPacket()
}
