package urlValidator

import (
	"fmt"
	"testing"
)

func TestURLValidate (t *testing.T) {
	var tests = []struct {
		url string
		valid bool
	}{
		{"google.com", true},
		{"www.google.com", true},
		{"wwww.google.com", true},
		{"salam.ir", true},
		{"sa1lam.ir", true},
		{"google", false},
		{"google@ir", false},
		{"google..com", false},
		{"goog1e.cOm", true},
	}

	for _, test := range tests {
		valid := Validate(test.url)
		if valid != test.valid {
			_ = fmt.Errorf("validate(%s) == %v, want %v", test.url, valid, test.valid)
		}
	}
}
