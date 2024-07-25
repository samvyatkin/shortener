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

var flags struct {
	serverAddr    *string
	shortenerAddr *string
}

func NewConfig() *Config {
	flag.StringVar(flags.serverAddr, "a", defaultPort, "server address")
	flag.StringVar(flags.shortenerAddr, "b", fmt.Sprintf("%s%s", defaultURL, defaultPort), "shortener address")
	flag.Parse()

	serverAddr := ""
	shortenerAddr := ""

	if addr, ok := os.LookupEnv(serverAddressKey); ok {
		serverAddr = addr
	}

	if addr, ok := os.LookupEnv(baseURLKey); ok {
		shortenerAddr = addr
	}

	if flags.serverAddr != nil && serverAddr == "" {
		serverAddr = *flags.serverAddr
	}

	if flags.shortenerAddr != nil && shortenerAddr == "" {
		shortenerAddr = *flags.shortenerAddr
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
