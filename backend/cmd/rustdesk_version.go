package cmd

import "strings"

func NormalizeRustdeskServerVersion(version string) string {
	v := strings.TrimSpace(version)
	if v == "" || v == "latest" {
		return "latest"
	}
	if strings.HasPrefix(v, "Branch_") {
		return strings.TrimPrefix(v, "Branch_")
	}
	return v
}
