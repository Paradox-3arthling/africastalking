package sms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/paradox-3arthling/africastalking"
)

func (req_data *Request_data) CheckInbox(last_received int) (map[string]interface{}, error) {
	inbox := make(map[string]interface{})
	if req_data.Prod == false {
		req_data.Username = "sandbox"
	}
	if req_data.Username == "" {
		return inbox, fmt.Errorf("`CheckInbox/1` found `req_data.Username` is blank!")
	}
	if req_data.Api_key == "" {
		return inbox, fmt.Errorf("`CheckInbox/1` found `req_data.Api_key` is blank!")
	}
	params := fmt.Sprintf("?username=%s&lastReceivedId=%v", req_data.Username, last_received)
	url := africastalking.SetUrl(req_data.Prod, africastalking.INBOX_URL) + params
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return inbox, fmt.Errorf("'http.NewRequest/3' got the error: %q", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("apiKey", req_data.Api_key)

	resp, err := client.Do(req)
	if err != nil {
		return inbox, fmt.Errorf("'client.Do/1' got the error: %q", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return inbox, fmt.Errorf("'ioutil.ReadAll/1' got the error: %q", err)
	}
	err = json.Unmarshal(body, &inbox)
	if err != nil {
		return inbox, fmt.Errorf("'ioutil.ReadAll/1' got the error: %q", err)
	}
	return inbox, nil
}
