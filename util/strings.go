package util

import "strings"

// ReplaceAllSlice uses strings.ReplaceAll on each index of the given slice. The slice IS modified, but also returned
func ReplaceAllSlice(in []string, old, new string) []string {
	for i, v := range in {
		in[i] = strings.ReplaceAll(v, old, new)
	}

	return in
}
