// package sms

// import "fmt"

// type Request_data map[string]string

// func (data Request_data) confirmFields() error {
// 	if data["username"] == "" {

// 	}

// 	return nil
// }

// func (data Request_data) SendSMS(prod bool) error {
// 	if prod == false {
// 		data["username"] = "sandbox"
// 	}
// 	if err := data.confirmFields(); err != nil {
// 		return fmt.Errorf("`SendSMS/1` got:\n%q", err)
// 	}

// 	return nil
// }
