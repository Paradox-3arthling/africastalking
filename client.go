package africastalking

import "strings"

func SetUrl(prod bool, url string) string {
	if prod {
		return strings.Replace(url, "sandbox.", "", 1)
	} else {
		return url
	}
}
