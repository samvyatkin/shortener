package config

import (
	"flag"
	"fmt"
	"os"
)

const (
	serverAddressKey = "SERVER_ADDRESS"
	baseURLKey       = "BASE_URL"
)

const (
	defaultURL  = "http://localhost"
	defaultPort = ":8080"
)

type Configuration interface {
	ServerAddr() string
	ShortenerAddr() string
}

type Config struct {
	serverAddr    string
	shortenerAddr string
}

type flags struct {
	serverAddr    string
	shortenerAddr string
}

func NewConfig() *Config {
	f := new(flags)
	flag.StringVar(&f.serverAddr, "a", defaultPort, "server address")
	flag.StringVar(&f.shortenerAddr, "b", fmt.Sprintf("%s%s", defaultURL, defaultPort), "shortener address")
	flag.Parse()

	serverAddr := ""
	shortenerAddr := ""

	if addr, ok := os.LookupEnv(serverAddressKey); ok {
		serverAddr = addr
	}

	if addr, ok := os.LookupEnv(baseURLKey); ok {
		shortenerAddr = addr
	}

	if serverAddr == "" {
		serverAddr = f.serverAddr
	}

	if shortenerAddr == "" {
		shortenerAddr = f.shortenerAddr
	}

	c := &Config{
		serverAddr:    serverAddr,
		shortenerAddr: shortenerAddr,
	}

	return c
}

func (c *Config) ServerAddr() string {
	return c.serverAddr
}

func (c *Config) ShortenerAddr() string {
	return c.shortenerAddr
}
