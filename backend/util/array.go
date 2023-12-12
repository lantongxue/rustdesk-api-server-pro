package util

func InStringArray(array []string, element string) bool {
	if array == nil || len(array) <= 0 || element == "" {
		return false
	}
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}
