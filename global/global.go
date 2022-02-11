package global

import (
	"github.com/spf13/viper"
	"github.com/the7s/swy-novel-server/config"
	"go.uber.org/zap"
)

var (
	SWY_CONFIG config.Server
	SWY_VP     *viper.Viper
	SWY_LOG    *zap.Logger
)
