package africastalking

import "testing"

// func TestPostClientCreated(t *testing.T) {
// 	client := BasicPostClient(SMS_URL)

// }

func TestSandboxOrProd(t *testing.T) {
	cases := []struct {
		prod      bool
		url, want string
	}{
		{false, SMS_URL, SMS_URL},
		{true, SMS_URL, "https://api.africastalking.com/version1/messaging"},
	}
	for i, val := range cases {
		got := SetUrl(val.prod, val.url)
		if got != val.want {
			t.Errorf("Test %v:\nSetUrl(%v, %q) == %q, want %q", i, val.prod, val.url, got, val.want)
		}
	}
}
