package cmd_test

import (
	"rustdesk-api-server-pro/cmd"
	"testing"
)

func TestRustdeskInstall_ShouldResolve146Release(t *testing.T) {
	tag := cmd.NormalizeRustdeskServerVersion("1.4.6")
	if tag != "1.4.6" {
		t.Fatalf("unexpected tag: %s", tag)
	}
}

func TestRustdeskInstall_ShouldResolveBranch146Release(t *testing.T) {
	tag := cmd.NormalizeRustdeskServerVersion("Branch_1.4.6")
	if tag != "1.4.6" {
		t.Fatalf("unexpected branch tag normalize: %s", tag)
	}
}
