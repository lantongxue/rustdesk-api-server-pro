package api_test

import (
	controller "rustdesk-api-server-pro/app/controller/api"
	"testing"
)

func TestAuthPayload_ShouldIgnoreUnknownAndKeepRequired(t *testing.T) {
	payload := map[string]any{
		"access_token": "x",
		"tfa_type":     "totp",
		"secret":       "s",
		"unknown":      "u",
	}
	normalized, ok := controller.NormalizeAuthPayload(payload)
	if !ok {
		t.Fatal("payload should be valid when required fields exist")
	}
	if normalized.AccessToken != "x" {
		t.Fatalf("unexpected access token: %s", normalized.AccessToken)
	}
}

func TestAuthPayload_ShouldRejectWhenAccessTokenMissing(t *testing.T) {
	payload := map[string]any{
		"tfa_type": "totp",
	}
	_, ok := controller.NormalizeAuthPayload(payload)
	if ok {
		t.Fatal("payload should be invalid without access token")
	}
}
