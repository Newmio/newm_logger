package logger

import (
	"context"
	"encoding/json"
	newm "newm/internal"

	"github.com/go-redis/redis/v8"
)

type ILoggerRepoRedis interface {
	CreateArrayLog(logs []Log) error
	CreateLog(log *Log) error
	GetLogs() ([]Log, error)
	GetLog() (Log, error)
}

type loggerRepoRedis struct {
	db *redis.Client
}

func NewLoggerRepoRedis(db *redis.Client) *loggerRepoRedis {
	return &loggerRepoRedis{db: db}
}

func (db *loggerRepoRedis) GetLog() (Log, error) {
	var log Log

	tx := db.db.TxPipeline()

	resultCmd := tx.LIndex(context.Background(), "logs", 0)

	_, err := tx.Exec(context.Background())
	if err != nil {
		return log, newm.Trace(err)
	}

	result, err := resultCmd.Result()
	if err != nil {
		return log, newm.Trace(err)
	}

	if result != "" {
		err := json.Unmarshal([]byte(result), &log)
		if err != nil {
			return log, newm.Trace(err)
		}

		err = db.db.LRem(context.Background(), "logs", 1, result).Err()
		if err != nil {
			return log, newm.Trace(err)
		}
	}

	return log, nil
}

func (db *loggerRepoRedis) GetLogs() ([]Log, error) {
	var logs []Log

	tx := db.db.TxPipeline()

	logsStr, err := tx.LRange(context.Background(), "logs", 0, -1).Result()
	if err != nil {
		return nil, newm.Trace(err)
	}

	for _, value := range logsStr {
		var log Log

		err = json.Unmarshal([]byte(value), &log)
		if err != nil {
			return nil, newm.Trace(err)
		}

		err := tx.LRem(context.Background(), "logs", 1, value).Err()
		if err != nil {
			return nil, newm.Trace(err)
		}

		logs = append(logs, log)
	}

	_, err = tx.Exec(context.Background())
	if err != nil {
		return nil, newm.Trace(err)
	}

	return logs, nil
}

func (db *loggerRepoRedis) CreateArrayLog(log []Log) error {
	body, err := json.Marshal(log)
	if err != nil {
		return newm.Trace(err)
	}

	err = db.db.RPush(context.Background(), "logs", body).Err()
	if err != nil {
		return newm.Trace(err)
	}

	return nil
}

func (db *loggerRepoRedis) CreateLog(log *Log) error {
	body, err := json.Marshal(log)
	if err != nil {
		return newm.Trace(err)
	}

	err = db.db.RPush(context.Background(), "logs", body).Err()
	if err != nil {
		return newm.Trace(err)
	}

	return nil
}
