package models

const MID_RRO_COM_INI = 0x0001
const MID_SRV_COM_INI = 0x0002

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

