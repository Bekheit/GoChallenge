package resources

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"os"
)

func NewLogger() (logger *zap.SugaredLogger, close func()) {
	level := zap.InfoLevel
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, level)
	log := zap.New(core, zap.AddCaller())

	closer := func() {
		_ = logger.Sync()
	}

	return log.Sugar(), closer
}
