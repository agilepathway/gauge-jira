// Package unique provides the functionality to create unique versions of slices.
package unique

// Strings returns a unique subset of the string slice provided.
func Strings(input []string) []string {
	uniques := make([]string, 0, len(input))
	keys := make(map[string]bool)

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (uniques)
	// then we append it.
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true

			uniques = append(uniques, entry)
		}
	}

	return uniques
}
