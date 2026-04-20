package version_test

import (
	"rustdesk-api-server-pro/helper/version"
	"testing"
)

func TestCapability_ShouldEnableTranslateMode_WhenVersionAtLeast146(t *testing.T) {
	capability := version.ResolveCapability("1.4.6")
	if !capability.TranslateMode {
		t.Fatal("translate mode should be enabled for 1.4.6")
	}
}

func TestCapability_ShouldDisableTranslateMode_WhenVersionLowerThan146(t *testing.T) {
	capability := version.ResolveCapability("1.4.5")
	if capability.TranslateMode {
		t.Fatal("translate mode should be disabled for versions below 1.4.6")
	}
}
