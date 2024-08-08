package utils

import "github.com/google/uuid"

type IdentifierGenerator interface {
	Generate() string
}

type UUIDGenerator struct{}

func (g *UUIDGenerator) Generate() string {
	return uuid.New().String()
}
