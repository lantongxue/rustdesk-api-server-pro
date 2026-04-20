package e2e

import (
	systemapi "rustdesk-api-server-pro/app/controller/api"
	"rustdesk-api-server-pro/helper/version"
	"testing"
)

type HandshakeResult struct {
	Success bool
}

func Run146HandshakeSmoke(t *testing.T) HandshakeResult {
	t.Helper()
	id := systemapi.ResolveHeartbeatRustdeskID("demo-id", "demo-uuid")
	normalized := systemapi.NormalizeReportedVersion("1.4.6", 0)
	capability := version.ResolveCapability(normalized)
	return HandshakeResult{
		Success: id != "" && capability.TranslateMode,
	}
}

func TestE2E_Smoke_146_ClientHandshake(t *testing.T) {
	result := Run146HandshakeSmoke(t)
	if !result.Success {
		t.Fatal("expected handshake smoke to succeed for 1.4.6")
	}
}
