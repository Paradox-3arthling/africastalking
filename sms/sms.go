package sms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/paradox-3arthling/africastalking"
)

const SMS_URL = "https://api.sandbox.africastalking.com/version1/messaging"

type Request_data struct {
	Username             string `json:"username"`
	To                   string `json:"to"`
	Message              string `json:"message"`
	From                 string `json:"from"`
	BulkSMSMode          string `json:"bulkSMSMode"`
	Enqueue              string `json:"enqueue"`
	Keyword              string `json:"keyword"`
	LinkId               string `json:"linkId"`
	RetryDurationInHours string `json:"retryDurationInHours"`
}

// func (req_data *Request_data) ConfirmFields() error {
// 	if req_data
// 	return nil
// }

// Return data at final func
func (req_data *Request_data) SendSMS(prod bool) error {
	data, err := json.Marshal(req_data)
	if err != nil {
		return fmt.Errorf("'json.Marshal/1' got the error: %q", err)
	}
	url := africastalking.SetUrl(prod, SMS_URL)
	req, err := africastalking.JsonRequest(url, "", data)
	if err != nil {
		return fmt.Errorf("'africastalking.JsonRequest/3' got the error: %q", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("'client.Do/1' got the error: %q", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("'ioutil.ReadAll/1' got the error: %q", err)
	}
	fmt.Printf("Body received: %q\n", string(body))
	return nil
}
