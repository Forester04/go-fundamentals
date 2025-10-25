package api_test

import (
	"backend/go/crypto/api"
	"testing"
)

func TestCextest(t *testing.T) {
	_, err := api.GetRateCurrency("Forester")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}
