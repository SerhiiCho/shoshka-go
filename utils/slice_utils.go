package utils

// Contains function returns true if slice contains given item
func Contains(slice []string, needle string) bool {
	for _, sliceItem := range slice {
		if sliceItem == needle {
			return true
		}
	}
	return false
}

// GetUniqueItem returns unique item from a slice
func GetUniqueItem(slice1 []string, slice2 []string) []string {
	result := []string{}

	for _, item1 := range slice1 {
		if !Contains(slice2, item1) {
			result = append(result, item1)
		}
	}

	return result
}
