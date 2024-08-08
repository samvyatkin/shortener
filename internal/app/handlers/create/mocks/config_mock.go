package mocks

type ConfigMock struct {
	serverAddr    string
	shortenerAddr string
}

func NewConfigMock(serverAddr, shortenerAddr string) *ConfigMock {
	return &ConfigMock{
		serverAddr:    serverAddr,
		shortenerAddr: shortenerAddr,
	}
}

func (c *ConfigMock) ServerAddr() string {
	return c.serverAddr
}

func (c *ConfigMock) ShortenerAddr() string {
	return c.shortenerAddr
}
