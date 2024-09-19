package utils

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	Logger = logger.Sugar()
}
