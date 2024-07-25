package config

import "flag"

type Configuration interface {
	ServerAddr() string
	ShortenerAddr() string
}

type AddrConfig struct {
	serverAddr    *string
	shortenerAddr *string
}

func NewAddrConfig() *AddrConfig {
	c := &AddrConfig{
		serverAddr:    flag.String("a", ":8080", "server address"),
		shortenerAddr: flag.String("b", "http://localhost:8080", "shortener address"),
	}

	flag.Parse()
	return c
}

func (c *AddrConfig) ServerAddr() string {
	return *c.serverAddr
}

func (c *AddrConfig) ShortenerAddr() string {
	return *c.shortenerAddr
}
