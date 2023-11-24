package logging

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() {
	var zapCfg zap.Config
	zapCfg = zap.NewProductionConfig()
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := zapCfg.Build()
	if err != nil {
		log.Fatalf("Failed to init logger %v", err)
	}

	zap.ReplaceGlobals(logger)
}
