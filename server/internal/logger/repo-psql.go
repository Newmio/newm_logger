package logger

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ILoggerRepoPsql interface {
	MigrateLogger() error
	CreateArrayLog(logs []Log) error
	CreateLog(log *Log) error
}

type loggerRepoPsql struct {
	db *sqlx.DB
}

func NewLoggerRepo(db *sqlx.DB) *loggerRepoPsql {
	return &loggerRepoPsql{db: db}
}

func (db *loggerRepoPsql) CreateLog(log *Log) error {
	str := `insert into logs(
		error, url, body_req, headers_req, status,
		body_resp, headers_resp, method, date_start,
		date_stop, milliseconds, ip, service, req_id, account_info) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10$, $11, $12, $13, $14, $15, $16)`

	_, err := db.db.Exec(str,
		log.Error, log.Url, log.BodyReq, log.HeadersReq,
		log.Status, log.BodyResp, log.HeadersResp, log.Method,
		log.DateStart, log.DateStop, log.Milliseconds, log.Ip,
		log.Service, log.RequestId, log.AccountInfo)
	if err != nil {
		return fmt.Errorf(err.Error() + "\n" + str)
	}

	return nil
}

func (db *loggerRepoPsql) CreateArrayLog(logs []Log) error {
	var str_values []string
	var values []interface{}

	for _, value := range logs {
		str_values = append(str_values,
			"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10$, $11, $12, $13, $14, $15, $16)")

		values = append(values,
			value.Error, value.Url, value.BodyReq, value.HeadersReq,
			value.Status, value.BodyResp, value.HeadersResp, value.Method,
			value.DateStart, value.DateStop, value.Milliseconds, value.Ip,
			value.Service, value.RequestId, value.AccountInfo)
	}

	str := `insert into logs(
		error, url, body_req, headers_req, status,
		body_resp, headers_resp, method, date_start,
		date_stop, milliseconds, ip, service,
		req_id, account_info) values`

	_, err := db.db.Exec(str+strings.Join(str_values, ","), values...)
	if err != nil {
		return fmt.Errorf(err.Error() + "\n" + str)
	}

	return nil
}
