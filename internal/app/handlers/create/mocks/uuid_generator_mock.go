package mocks

type UUIDGeneratorMock struct {
	uuid string
}

func NewUUIDGeneratorMock(uuid string) *UUIDGeneratorMock {
	return &UUIDGeneratorMock{uuid: uuid}
}

func (g *UUIDGeneratorMock) Generate() string {
	return g.uuid
}
