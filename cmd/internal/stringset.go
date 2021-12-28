package internal

import "sort"

// A string set which maintains an ascending sort order.
type StringSet []string

// Insert adds a new string to the sorted string set if it is not already
// present.
func (s StringSet) Insert(newString string) StringSet {
	index := sort.SearchStrings(s, newString)
	if index < len(s) && s[index] == newString {
		// Don't need to insert if the value is already present in the set.
		return s
	}
	s = append(s, "")
	copy(s[index+1:], s[index:])
	s[index] = newString
	return s
}
