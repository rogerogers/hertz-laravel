package utils

// InArray check if a string in a slice
func InArray(key string, arr []string) bool {
	for _, k := range arr {
		if k == key {
			return true
		}
	}
	return false
}
