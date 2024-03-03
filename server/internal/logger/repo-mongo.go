package logger

import "go.mongodb.org/mongo-driver/mongo"

type ILoggerRepoMongo interface {
	// MigrateLogger() error
	// CreateArrayLog(logs []Log) error
	// CreateLog(log *Log) error
}

type loggerRepoMongo struct {
	db *mongo.Database
}

func NewLoggerRepoMongo(db *mongo.Database)*loggerRepoMongo{
	return &loggerRepoMongo{db: db}
}