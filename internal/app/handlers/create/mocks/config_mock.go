package mocks

type ConfigMock struct {
	serverAddr      string
	shortenerAddr   string
	fileStoragePath string
}

func NewConfigMock(serverAddr, shortenerAddr, fileStoragePath string) *ConfigMock {
	return &ConfigMock{
		serverAddr:      serverAddr,
		shortenerAddr:   shortenerAddr,
		fileStoragePath: fileStoragePath,
	}
}

func (c *ConfigMock) ServerAddr() string {
	return c.serverAddr
}

func (c *ConfigMock) ShortenerAddr() string {
	return c.shortenerAddr
}

func (c *ConfigMock) FileStoragePath() string {
	return c.fileStoragePath
}
