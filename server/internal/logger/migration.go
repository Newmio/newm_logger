package logger

import "fmt"

func (db *loggerRepoPsql) MigrateLogger() error {
	str := `create table if not exists logs(
		id primary key,
		error text default '',
		url text default '',
		body_req text default '',
		headers_req text default '',
		status int default 0,
		body_resp text default '',
		headers_resp text default '',
		method text default '',
		date_start timestamp default '2024-01-01 08:08:08.000',
		date_stop timestamp default '2024-01-01 08:08:08.000',
		milliseconds int default 0,
		ip text default '',
		service text default '',
		req_id text default '',
		account_info text default ''
	)`

	_, err := db.db.Exec(str)
	if err != nil {
		return fmt.Errorf(err.Error() + "\n" + str)
	}

	return nil
}
