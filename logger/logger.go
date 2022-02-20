package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},                  // loggea en el standard output
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel), // info level and above
		Encoding:    "json",                              // all in json format
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level", // key level -> level of log Ex: "level": "info"
			TimeKey:      "time",  // time key Ex: "time": "2029-11-10T10:10:10"
			MessageKey:   "msg",   // message key
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	log.Sync()
}
