package config

type Configuration interface {
	ServerAddr() string
	ShortenerAddr() string
}
