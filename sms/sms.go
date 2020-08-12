package sms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/paradox-3arthling/africastalking"
)

// check how `json.Marshal/1` marshals arrays
type Request_data struct {
	Username             string   `json:"username,omitempty"`
	PhoneNumbers         []string `json:"phoneNumbers,omitempty"`
	Message              string   `json:"message,omitempty"`
	From                 string   `json:"senderId,omitempty"`
	BulkSMSMode          int      `json:"bulkSMSMode,omitempty"`
	Enqueue              int      `json:"enqueue,omitempty"`
	Keyword              string   `json:"keyword,omitempty"`
	LinkId               string   `json:"linkId,omitempty"`
	RetryDurationInHours int      `json:"retryDurationInHours,omitempty"`
}

// Need to confirm that the nos. are valid
func (req_data *Request_data) ConfirmFields() error {
	if req_data.Username == "" {
		return fmt.Errorf("`ConfirmFields` got `req_data.Username` is blank!")
	}
	if len(req_data.PhoneNumbers) == 0 {
		return fmt.Errorf("`ConfirmFields` got `req_data.To` has no numbers to send to!")
	}
	if req_data.Message == "" {
		return fmt.Errorf("`ConfirmFields` got `req_data.Message` is blank!")
	}
	return nil
}

func (req_data *Request_data) encodeValues() string {
	data := url.Values{}
	data.Set("", "")

	return data.Encode()
}
func setValue(data *url.Values, value string) {

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
	url := africastalking.SetUrl(prod, africastalking.SMS_URL)
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
	fmt.Printf("Body received: %q\nURL: %q\ndata: %q", string(body), url, string(data))
	return nil
}
