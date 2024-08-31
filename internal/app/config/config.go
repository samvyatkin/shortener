package config

import (
	"flag"
	"fmt"
	"os"
)

const (
	serverAddressKey   = "SERVER_ADDRESS"
	baseURLKey         = "BASE_URL"
	fileStoragePathKey = "FILE_STORAGE_PATH"
)

const (
	defaultURL             = "http://localhost"
	defaultPort            = ":8080"
	defaultFileStoragePath = "data.txt"
)

type Configuration interface {
	ServerAddr() string
	ShortenerAddr() string
	FileStoragePath() string
}

type Config struct {
	serverAddr      string
	shortenerAddr   string
	fileStoragePath string
}

type flags struct {
	serverAddr      string
	shortenerAddr   string
	fileStoragePath string
}

func NewConfig() *Config {
	f := new(flags)
	flag.StringVar(&f.serverAddr, "a", defaultPort, "server address")
	flag.StringVar(&f.shortenerAddr, "b", fmt.Sprintf("%s%s", defaultURL, defaultPort), "shortener address")
	flag.StringVar(&f.fileStoragePath, "f", defaultFileStoragePath, "file storage path")
	flag.Parse()

	serverAddr := ""
	shortenerAddr := ""
	fileStoragePath := ""

	if addr, ok := os.LookupEnv(serverAddressKey); ok {
		serverAddr = addr
	}

	if addr, ok := os.LookupEnv(baseURLKey); ok {
		shortenerAddr = addr
	}

	if path, ok := os.LookupEnv(fileStoragePathKey); ok {
		fileStoragePath = path
	}

	if serverAddr == "" {
		serverAddr = f.serverAddr
	}

	if shortenerAddr == "" {
		shortenerAddr = f.shortenerAddr
	}

	if fileStoragePath == "" {
		fileStoragePath = f.fileStoragePath
	}

	c := &Config{
		serverAddr:      serverAddr,
		shortenerAddr:   shortenerAddr,
		fileStoragePath: fileStoragePath,
	}

	return c
}

func (c *Config) ServerAddr() string {
	return c.serverAddr
}

func (c *Config) ShortenerAddr() string {
	return c.shortenerAddr
}

func (c *Config) FileStoragePath() string {
	return c.fileStoragePath
}
