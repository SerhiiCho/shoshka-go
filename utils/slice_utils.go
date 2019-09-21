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
	var result []string

	for _, item1 := range slice1 {
		if !Contains(slice2, item1) {
			result = append(result, item1)
		}
	}

	for _, item2 := range slice2 {
		if !Contains(slice1, item2) {
			result = append(result, item2)
		}
	}

	return result
}

// GetIndexOfSliceItem returns index of a slice item where value is positioned
func GetIndexOfSliceItem(slice []string, value string) int {
	for i, sliceItem := range slice {
		if sliceItem == value {
			return i
		}
	}

	return -1
}
