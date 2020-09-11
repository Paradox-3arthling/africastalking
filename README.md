![Go](https://github.com/Paradox-3arthling/africastalking/workflows/Go/badge.svg)
# africastalking
## What it does
This is an Africastalking API wrapper in Golang.

> This is a project is still in dev, though working parts
> should be ok.

## Why is it useful
- Trying to make it easier to use Africastalking API so as developers can concentrate on development.
- Could better help people learn more about golang 'net/http' package as it heavily uses this package.
- Also, hope to make very fast application(hope so :P compiled source should be faster) with the beauty of concurrency.

## How to get started
> !!IMPORTANT!! do check for errors, am ignoring them for the sake of short examples.
This will be sandbox examples:
### SMS package
The import statement for this will be `import "github.com/paradox-3arthling/africastalking/sms"`. Which has a struct
`Request_data{}` which can be use to make 2 request:
1. Sending bulk SMS's
  ```
	req_data := sms.Request_data{
		Api_key: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		To:      []string{"+254XXXXXXXXX", "+254XXXXXXXXX"},
		Message: "SMS imeingiana",
	}
	map_resp, _ := req_data.SendSMS()
    // other field is an error if found.
	fmt.Printf("Return map: %v", map_resp)
  ```
2. Fetch inbox messages
  ```
	req_data := sms.Request_data{
		Api_key: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}
	map_resp, _ := req_data.CheckInbox(48739)
	// other field is an error if found.
	fmt.Printf("Return map: %v", map_resp)
  ```

## WIP
- Check issues for next feature updates to be done.

## Who maintains and can contributes.
At the moment am the only maintainer and contributer open to help, contact me at floydqaranja@gmail.com 
