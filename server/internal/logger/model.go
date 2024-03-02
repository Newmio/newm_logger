package logger

type Log struct {
	Id           int    `db:"id" json:"id" xml:"id"`
	Error        string `db:"error" json:"error" xml:"error"`
	Url          string `db:"url" json:"url" xml:"url"`
	BodyReq      string `db:"body_req" json:"body_req" xml:"body_req"`
	HeadersReq   string `db:"headers_req" json:"headers_req" xml:"headers_req"`
	Status       int    `db:"status" json:"status" xml:"status"`
	BodyResp     string `db:"body_resp" json:"body_resp" xml:"body_resp"`
	HeadersResp  string `db:"headers_resp" json:"headers_resp" xml:"headers_resp"`
	Method       string `db:"method" json:"method" xml:"method"`
	DateStart    string `db:"date_start" json:"date_start" xml:"date_start"`
	DateStop     string `db:"date_stop" json:"date_stop" xml:"date_stop"`
	Milliseconds int    `db:"milliseconds" json:"milliseconds" xml:"milliseconds"`
	Ip           string `db:"ip" json:"ip" xml:"ip"`
	Service      string `db:"service" json:"service" xml:"service"`
	RequestId    string `db:"req_id" json:"req_id" xml:"req_id"`
	AccountInfo  string `db:"account_info" json:"account_info" xml:"account_info"`
}

func errorRespnse(err error) map[string]interface{} {
	return map[string]interface{}{"error": err.Error()}
}
