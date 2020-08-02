package africastalking

import "testing"

const SMS_URL = "https://api.sandbox.africastalking.com/version1/messaging"

// func TestPostClientCreated(t *testing.T) {
// 	client := BasicPostClient(SMS_URL)

// }

func TestSandboxOrProd(t *testing.T) {
	cases := []struct {
		prod      bool
		url, want string
	}{
		{false, SMS_URL, "https://api.sandbox.africastalking.com/version1/messaging"},
		{true, SMS_URL, "https://api.africastalking.com/version1/messaging"},
	}
	for _, val := range cases {
		got := SetUrl(val.prod, val.url)
		if got != val.want {
			t.Errorf("SetUrl(%v, %q) == %q, want %q", val.prod, val.url, got, val.want)
		}
	}
}
