package statuses

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func parseJSONFromFile() map[int]string {
	data, err := ioutil.ReadFile("./statuses.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var statuses map[int]string
	json.Unmarshal(data, &statuses)

	return statuses
}

// Statuses contains all statuses code and messages
var Statuses = parseJSONFromFile()
var emptyCode = []int{204, 205, 305}
var retryCode = []int{502, 503, 504}

// RedirectCode check a status code, return if it's a redirect code
func RedirectCode(code int) bool {
	if code <= 308 && code >= 300 {
		return true
	}

	return false
}

// EmptyCode check a status code, return if it's a empty code
func EmptyCode(code int) bool {
	for _, v := range emptyCode {
		if v == code {
			return true
		}
	}

	return false
}

// RetryCode check a status code, return if it's a retry code
func RetryCode(code int) bool {
	for _, v := range retryCode {
		if v == code {
			return true
		}
	}

	return false
}

// Status returns the message of the specified code
func Status(code int) string {
	return Statuses[code]
}
