package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

func NewZapLog(filename string) *ZapLog {
	rawJSON := []byte(`{
		"level": "info",
		"outputPaths": ["` + filename + `"],
		"errorOutputPaths": ["stderr"],
		"encoding": "json",
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		}
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return &ZapLog{
		log: log.Sugar(),
	}
}

type ZapLog struct {
	log *zap.SugaredLogger
}

func (l *ZapLog) Info(v ...interface{}) {
	l.log.Info("msg", v)
}
