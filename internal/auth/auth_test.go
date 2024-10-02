package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey unit_test")
	key, err := GetAPIKey(header)
	if err != nil || key == "" {
		t.Fail()
	}
}
