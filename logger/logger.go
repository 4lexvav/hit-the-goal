package logger

import (
	lgr "github.com/eugeneradionov/logger"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Get() *zap.SugaredLogger {
	return logger.Sugar()
}

func Load(preset string) (err error) {
	logger, err = lgr.Load(lgr.LogPreset((preset)))
	return err
}
