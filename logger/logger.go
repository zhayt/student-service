package logger

import (
	"github.com/zhayt/student-service/config"
	"go.uber.org/zap"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	if cfg.AppMode == "Dev" {
		return zap.NewDevelopment()
	} else {
		return zap.NewProduction()
	}
}
