package api_test

import (
	systemapi "rustdesk-api-server-pro/app/controller/api"
	"testing"
)

func TestPostHeartbeat_ShouldAccept_1_4_6Payload(t *testing.T) {
	id := systemapi.ResolveHeartbeatRustdeskID("test", "u-1")
	if id != "test" {
		t.Fatalf("unexpected rustdesk id: %s", id)
	}
}

func TestPostHeartbeat_ShouldFallbackToUUIDWhenIDMissing(t *testing.T) {
	id := systemapi.ResolveHeartbeatRustdeskID("", "u-1")
	if id != "u-1" {
		t.Fatalf("unexpected fallback rustdesk id: %s", id)
	}
}

func TestNormalizeReportedVersion_ShouldUseNumericVersionWhenStringMissing(t *testing.T) {
	version := systemapi.NormalizeReportedVersion("", 10604)
	if version != "10604" {
		t.Fatalf("unexpected normalized version: %s", version)
	}
}
