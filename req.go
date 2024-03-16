package newm_logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//////////

func CreateLog(log Log)error{
	body, err := json.Marshal(log)
	if err != nil{
		return err
	}

	return req(body)
}

func CreateArrayLog(log []Log)error{
	body, err := json.Marshal(log)
	if err != nil{
		return err
	}

	return req(body)
}

func req(body []byte) error {
	client := &http.Client{}

	req, err := http.NewRequest("POST", URL_LOGGER, bytes.NewBuffer(body))
	if err != nil{
		return err
	}

	resp, err := client.Do(req)
	if err != nil{
		return err
	}
	defer resp.Body.Close()

	var respLog respLog

	body, err = io.ReadAll(resp.Body)
	if err != nil{
		return err
	}

	err = json.Unmarshal(body, &respLog)
	if err != nil{
		return err
	}

	if respLog.Status != "ok"{
		return fmt.Errorf(respLog.Error)
	}

	return nil
}
