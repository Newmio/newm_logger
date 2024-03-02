package logger

type ILoggerService interface {
	CreateLog(log *Log) error
	CreateArrayLog(logs []Log) error
}

type loggerService struct {
	psql ILoggerRepoPsql
}

func NewLoggerService(psql ILoggerRepoPsql) *loggerService {
	err := psql.MigrateLogger()
	if err != nil {
		return nil
	}
	return &loggerService{psql: psql}
}

func (s *loggerService) CreateLog(log *Log) error {
	return s.psql.CreateLog(log)
}

func (s *loggerService) CreateArrayLog(logs []Log) error {
	if len(logs) > 0 {
		return s.psql.CreateArrayLog(logs)
	}

	return nil
}
