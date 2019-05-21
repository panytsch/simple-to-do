package helpers

//ContainsString - check is item in slice
func ContainsString(slice []string, element string) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}
	return false
}
