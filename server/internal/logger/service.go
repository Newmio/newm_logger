package logger

import "time"

type ILoggerService interface {
	CreateLog(log *Log) error
	CreateArrayLog(logs []Log) error
}

type loggerService struct {
	psql ILoggerRepoPsql
	mongo ILoggerRepoMongo
	redis ILoggerRepoRedis
}

func NewLoggerService(psql ILoggerRepoPsql, mongo ILoggerRepoMongo, redis ILoggerRepoRedis) *loggerService {
	err := psql.MigrateLogger()
	if err != nil {
		return nil
	}

	service := &loggerService{psql: psql, mongo: mongo, redis: redis}

	service.tickerForSetLogsInDb()

	return service
}

func (s *loggerService) CreateLog(log *Log) error {
	return s.redis.CreateLog(log)
}

func (s *loggerService) CreateArrayLog(logs []Log) error {
	if len(logs) > 0 {
		return s.redis.CreateArrayLog(logs)
	}

	return nil
}

func (s *loggerService) tickerForSetLogsInDb(){
	ticker := time.NewTicker(100 * time.Millisecond)

	for range ticker.C{

		log, err := s.redis.GetLog()
		if err != nil{
			continue
		}

		err = s.psql.CreateLog(&log)
		if err != nil{
			continue
		}
	}
}

