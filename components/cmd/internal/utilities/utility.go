package utilities

// StringInArray checks if a string is in the array
// It returns a bool value
func StringInArray(a string, arr []string) bool {
	for _, b := range arr {
		if b == a {
			return true
		}
	}
	return false
}
