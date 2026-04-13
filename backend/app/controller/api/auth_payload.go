package api

import "strings"

type AuthPayload struct {
	AccessToken string
	TfaType     string
	Secret      string
}

func NormalizeAuthPayload(payload map[string]any) (AuthPayload, bool) {
	result := AuthPayload{
		AccessToken: toString(payload["access_token"]),
		TfaType:     toString(payload["tfa_type"]),
		Secret:      toString(payload["secret"]),
	}

	if result.AccessToken == "" {
		return result, false
	}
	return result, true
}

func toString(v any) string {
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(s)
}
