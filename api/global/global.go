package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/minibear2333/programmer-go/api/internal/config"
	"go.uber.org/zap"
)

var (
	LOG    *zap.Logger
	REDIS  *redis.Client
	CONFIG config.Config
)
