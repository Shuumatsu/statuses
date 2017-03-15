package statuses

import (
	"testing"
)

func TestRetryCode(t *testing.T) {
	codes := []int{502, 503, 504}
	for _, v := range codes {
		if !RetryCode(v) {
			t.Errorf("%d code not successfully validadted", v)
		}
	}
}

func TestEmpty(t *testing.T) {
	codes := []int{204, 205, 305}
	for _, v := range codes {
		if !EmptyCode(v) {
			t.Errorf("%d code not successfully validadted", v)
		}
	}
}

func TestRedirectCode(t *testing.T) {
	codes := make([]int, 9)
	for i := 0; i <= 8; i++ {
		codes[i] = 300 + i
	}

	for _, v := range codes {
		if !RedirectCode(v) {
			t.Errorf("%d code not successfully validadted", v)
		}
	}
}

func TestStatus(t *testing.T) {
	invalidCode := 9999
	if Status(invalidCode) != "" {
		t.Errorf("Status does not return an empty string")
	}

	if Status(302) != "Found" {
		t.Errorf("Status does not return the correct message")
	}
}
