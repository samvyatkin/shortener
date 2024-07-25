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

type FlagConfig struct {
	serverAddr    string
	shortenerAddr string
}

func NewFlagConfig() *FlagConfig {
	serverAddr := ""
	shortenerAddr := ""

	if addr, ok := os.LookupEnv(serverAddressKey); ok {
		serverAddr = addr
	}

	if addr, ok := os.LookupEnv(baseURLKey); ok {
		shortenerAddr = addr
	}

	if serverAddr == "" {
		flag.StringVar(&serverAddr, "a", defaultPort, "server address")
	}

	if shortenerAddr == "" {
		flag.StringVar(&shortenerAddr, "b", fmt.Sprintf("%s%s", defaultURL, defaultURL), "shortener address")
	}

	c := &FlagConfig{
		serverAddr:    serverAddr,
		shortenerAddr: shortenerAddr,
	}

	flag.Parse()
	return c
}

func (c *FlagConfig) ServerAddr() string {
	return c.serverAddr
}

func (c *FlagConfig) ShortenerAddr() string {
	return c.shortenerAddr
}
