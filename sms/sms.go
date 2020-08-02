package sms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/paradox-3arthling/africastalking"
)

const SMS_URL = "https://api.sandbox.africastalking.com/version1/messaging"

// check how `json.Marshal/1` marshals arrays
type Request_data struct {
	Username             string   `json:"username"`
	To                   []string `json:"to"`
	Message              string   `json:"message"`
	From                 string   `json:"from"`
	BulkSMSMode          string   `json:"bulkSMSMode"`
	Enqueue              string   `json:"enqueue"`
	Keyword              string   `json:"keyword"`
	LinkId               string   `json:"linkId"`
	RetryDurationInHours string   `json:"retryDurationInHours"`
}

// Need to confirm that the nos. are valid
func (req_data *Request_data) ConfirmFields() error {
	if req_data.Username == "" {
		return fmt.Errorf("`ConfirmFields` got `req_data.Username` is blank!")
	}
	if len(req_data.To) == 0 {
		return fmt.Errorf("`ConfirmFields` got `req_data.To` has no numbers to send to!")
	}
	if req_data.Message == "" {
		return fmt.Errorf("`ConfirmFields` got `req_data.Message` is blank!")
	}
	return nil
}

// Return data at final func
func (req_data *Request_data) SendSMS(prod bool) error {
	if prod == false {
		req_data.Username = "sandbox"
	}
	err := req_data.ConfirmFields()
	if err != nil {
		return fmt.Errorf("'req_data.ConfirmFields/0' got the error: %q", err)
	}
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
