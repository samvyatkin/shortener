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
	dbAddrKey          = "DATABASE_DSN"
)

const (
	defaultURL             = "http://localhost"
	defaultPort            = ":8080"
	defaultFileStoragePath = "data.txt"
	defaultDBAddress       = ""
)

type Configuration interface {
	ServerAddr() string
	ShortenerAddr() string
	FileStoragePath() string
	DBAddr() string
}

type Config struct {
	serverAddr      string
	shortenerAddr   string
	fileStoragePath string
	dbAddr          string
}

type flags struct {
	serverAddr      string
	shortenerAddr   string
	fileStoragePath string
	dbAddr          string
}

func NewConfig() *Config {
	f := new(flags)
	flag.StringVar(&f.serverAddr, "a", defaultPort, "server address")
	flag.StringVar(&f.shortenerAddr, "b", fmt.Sprintf("%s%s", defaultURL, defaultPort), "shortener address")
	flag.StringVar(&f.fileStoragePath, "f", defaultFileStoragePath, "file storage path")
	flag.StringVar(&f.dbAddr, "d", defaultDBAddress, "database address")
	flag.Parse()

	serverAddr := ""
	shortenerAddr := ""
	fileStoragePath := ""
	dbAddr := ""

	if addr, ok := os.LookupEnv(serverAddressKey); ok {
		serverAddr = addr
	}

	if addr, ok := os.LookupEnv(baseURLKey); ok {
		shortenerAddr = addr
	}

	if path, ok := os.LookupEnv(fileStoragePathKey); ok {
		fileStoragePath = path
	}

	if addr, ok := os.LookupEnv(dbAddrKey); ok {
		dbAddr = addr
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

	if dbAddr == "" {
		dbAddr = f.dbAddr
	}

	c := &Config{
		serverAddr:      serverAddr,
		shortenerAddr:   shortenerAddr,
		fileStoragePath: fileStoragePath,
		dbAddr:          dbAddr,
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

func (c *Config) DBAddr() string {
	return c.dbAddr
}
