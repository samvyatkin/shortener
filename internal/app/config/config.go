package config

import "flag"

type Configuration struct {
	ServerAddr    *string
	ShortenerAddr *string
}

func New() *Configuration {
	c := &Configuration{
		ServerAddr:    flag.String("a", ":8080", "server address"),
		ShortenerAddr: flag.String("b", "localhost:8080", "shortener address"),
	}

	flag.Parse()
	return c
}
