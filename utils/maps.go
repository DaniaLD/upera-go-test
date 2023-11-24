package utils

func MapsDiff(oldMap, newMap map[string]string) (updatedAttributes []string) {
	for key, oldValue := range oldMap {
		// Check if the key exists in the new map
		if newValue, ok := newMap[key]; ok {
			// Compare the values
			if newValue != oldValue {
				updatedAttributes = append(updatedAttributes, key)
			}
		}
	}

	return updatedAttributes
}
