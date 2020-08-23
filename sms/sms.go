package sms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/paradox-3arthling/africastalking"
)

// check how `json.Marshal/1` marshals arrays
type Request_data struct {
	Prod                 bool
	Api_key              string
	Username             string   //`username`
	To                   []string //`to`
	Message              string   //`message`
	From                 string   //`from`
	BulkSMSMode          int      //`bulkSMSMode`
	Enqueue              int      //`enqueue`
	Keyword              string   //`keyword`
	LinkId               string   //`linkId`
	RetryDurationInHours int      //`retryDurationInHours`
}

func (req_data *Request_data) encodeValues() []byte {
	data := url.Values{}
	data.Set("username", req_data.Username)
	data.Set("to", strings.Join(req_data.To, ","))
	data.Set("message", req_data.Message)
	setDataStr(&data, req_data.From, "from")
	setDataInt(&data, req_data.BulkSMSMode, "bulkSMSMode")
	setDataInt(&data, req_data.Enqueue, "enqueue")
	setDataStr(&data, req_data.Keyword, "keyword")
	setDataStr(&data, req_data.LinkId, "linkId")
	setDataInt(&data, req_data.RetryDurationInHours, "retryDurationInHours")

	return []byte(data.Encode())
}
func setDataStr(data *url.Values, val, key string) {
	if val != "" {
		data.Set(key, val)
	}
}
func setDataInt(data *url.Values, val int, key string) {
	if val != 0 {
		data.Set(key, string(val))
	}
}

// Need to confirm that the nos. are valid
func (req_data *Request_data) ConfirmFields() error {
	if req_data.Api_key == "" {
		return fmt.Errorf("`ConfirmFields` got `req_data.Api_key` is blank!")
	}
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
func (req_data *Request_data) SendSMS() (map[string]interface{}, error) {
	json_map := make(map[string]interface{})

	prod := req_data.Prod
	if prod == false {
		req_data.Username = "sandbox"
	}
	err := req_data.ConfirmFields()
	if err != nil {
		return json_map, fmt.Errorf("'req_data.ConfirmFields/0' got the error: %q", err)
	}
	data := req_data.encodeValues()
	url := africastalking.SetUrl(prod, africastalking.SMS_URL)
	req, err := africastalking.EncodedRequest(url, req_data.Api_key, data)
	if err != nil {
		return json_map, fmt.Errorf("'africastalking.JsonRequest/3' got the error: %q", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return json_map, fmt.Errorf("'client.Do/1' got the error: %q", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return json_map, fmt.Errorf("'ioutil.ReadAll/1' got the error: %q", err)
	}
	// fmt.Printf("URL: %q\ndata: %q\n", url, string(data))
	err = json.Unmarshal(body, &json_map)
	if err != nil {
		return json_map, fmt.Errorf("'json.Unmarshal/1' got the error: %q", err)
	}
	return json_map, nil
}
