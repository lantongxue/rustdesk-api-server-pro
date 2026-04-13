package version

import (
	"strconv"
	"strings"
)

type Capability struct {
	TranslateMode bool `json:"translate_mode"`
}

func ResolveCapability(version string) Capability {
	return Capability{
		TranslateMode: versionGTE(version, "1.4.6"),
	}
}

func versionGTE(version string, target string) bool {
	v := parseVersion(version)
	t := parseVersion(target)
	for i := 0; i < 3; i++ {
		if v[i] > t[i] {
			return true
		}
		if v[i] < t[i] {
			return false
		}
	}
	return true
}

func parseVersion(version string) [3]int {
	version = strings.TrimSpace(strings.TrimPrefix(version, "v"))
	parts := strings.Split(version, ".")
	parsed := [3]int{}
	for i := 0; i < len(parts) && i < 3; i++ {
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			return [3]int{}
		}
		parsed[i] = n
	}
	return parsed
}
