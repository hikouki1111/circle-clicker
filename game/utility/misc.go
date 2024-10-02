package utility

import (
	"fmt"
	"strings"
	"syscall/js"
)

func ParseCookie(document js.Value) map[string]string {
	cookieStr := document.Get("cookie").String()
	cookies := map[string]string{}

	if cookieStr == "" {
		return nil
	}

	cookieArray := strings.Split(cookieStr, "; ")
	for _, cookie := range cookieArray {
		pair := strings.SplitN(cookie, "=", 2)
		if len(pair) == 2 {
			cookies[pair[0]] = pair[1]
		}
	}

	return cookies
}

func ParseBool(boolStr string) (bool, error) {
	boolStr = strings.ToLower(boolStr)
	if boolStr == "true" {
		return true, nil
	} else if boolStr == "false" {
		return false, nil
	}

	return false, fmt.Errorf("error")
}
