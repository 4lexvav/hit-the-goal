package logger

import (
	"context"

	reqContext "github.com/4lexvav/hit-the-goal/context"
	lgr "github.com/eugeneradionov/logger"
	"go.uber.org/zap"
)

const requestIDKey = "request_id"

var logger *zap.Logger

func Get() *zap.SugaredLogger {
	return logger.Sugar()
}

func Load(preset string) (err error) {
	logger, err = lgr.Load(lgr.LogPreset((preset)))
	return err
}

func WithCtxValue(ctx context.Context) *zap.SugaredLogger {
	return logger.Sugar().With(requestIDKey, reqContext.GetRequestID(ctx))
}
