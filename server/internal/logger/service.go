package logger

type ILoggerService interface{
	CreateLog(log *Log)error
	CreateArrayLog(logs []Log)error
}

type loggerService struct{
	r ILoggerRepo
}

func NewLoggerService(r ILoggerRepo)*loggerService{
	err := r.MigrateLogger()
	if err != nil{
		return nil
	}
	return &loggerService{r: r}
}

func (s *loggerService) CreateLog(log *Log)error{
	return s.r.CreateLog(log)
}

func (s *loggerService) CreateArrayLog(logs []Log)error{
	if len(logs) > 0{
		return s.r.CreateArrayLog(logs)
	}

	return nil
}