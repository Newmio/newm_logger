package logger

import "github.com/go-redis/redis/v8"

type ILoggerRepoRedis interface {
	// MigrateLogger() error
	// CreateArrayLog(logs []Log) error
	// CreateLog(log *Log) error
}

type loggerRepoRedis struct {
	db *redis.Client
}

func NewLoggerRepoRedis(db *redis.Client)*loggerRepoRedis{
	return &loggerRepoRedis{db: db}
}