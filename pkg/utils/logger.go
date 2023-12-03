package utils

import (
	zap "go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
