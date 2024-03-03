package logger

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type ILoggerRepoRedis interface {
	CreateArrayLog(logs []Log) error
	CreateLog(log *Log) error
	GetLogs() ([]Log, error)
	GetLog()(Log, error)
}

type loggerRepoRedis struct {
	db *redis.Client
}

func NewLoggerRepoRedis(db *redis.Client) *loggerRepoRedis {
	return &loggerRepoRedis{db: db}
}

func (db *loggerRepoRedis) GetLog()(Log, error){
	var log Log

	tx := db.db.TxPipeline()

	result, err := tx.LIndex(context.Background(), "logs", 0).Result()
	if err != nil{
		return log, err
	}

	err = json.Unmarshal([]byte(result), &log)
	if err != nil{
		return log, err
	}

	err = tx.LRem(context.Background(), "logs", 1, result).Err()
	if err != nil{
		return log, err
	}

	_, err = tx.Exec(context.Background())
	if err != nil{
		return log, err
	}

	return log, nil
}

func (db *loggerRepoRedis) GetLogs() ([]Log, error) {
	var logs []Log

	tx := db.db.TxPipeline()

	logsStr, err := tx.LRange(context.Background(), "logs", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	for _, value := range logsStr {
		var log Log

		err = json.Unmarshal([]byte(value), &log)
		if err != nil {
			return nil, err
		}

		err := tx.LRem(context.Background(), "logs", 1, value).Err()
		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	_, err = tx.Exec(context.Background())
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (db *loggerRepoRedis) CreateArrayLog(log []Log) error {
	err := db.db.RPush(context.Background(), "logs", log).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *loggerRepoRedis) CreateLog(log *Log) error {
	err := db.db.RPush(context.Background(), "logs", log).Err()
	if err != nil {
		return err
	}

	return nil
}
