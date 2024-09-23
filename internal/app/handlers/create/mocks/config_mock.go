package mocks

type ConfigMock struct {
	serverAddr      string
	shortenerAddr   string
	fileStoragePath string
	dbAddr          string
}

func NewConfigMock(serverAddr, shortenerAddr, fileStoragePath, dbAddr string) *ConfigMock {
	return &ConfigMock{
		serverAddr:      serverAddr,
		shortenerAddr:   shortenerAddr,
		fileStoragePath: fileStoragePath,
		dbAddr:          dbAddr,
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

func (c *ConfigMock) DBAddr() string {
	return c.dbAddr
}
